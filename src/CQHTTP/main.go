// udp_client project main.go
package main

import (
	"HMsg"
	"cqencode"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

//Recive boardcast from CQS server via UDP
func RcvMsgFromCQ(port int) {
	socket, e := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: port,
	})

	defer socket.Close()

	if e != nil {
		log.Println(e)
		return
	} else {
		log.Println("Receive msg via: ", port)
	}

	for {
		log.Println("waiting message...")
		data := make([]byte, PreFixMaxSize+PayLoadMaxSize)

		n, _, _ := socket.ReadFromUDP(data[0:])

		dataStr := string(data[0:n])

		log.Printf("Origin message: %s", dataStr)

		frameParts := strings.Split(dataStr, " ")
		EncodedText := frameParts[MapMsgIdx[frameParts[0]].EncodedText]
		if EncodedText != frameParts[0] {
			gbk := cqencode.Base64_gbk(EncodedText)
			log.Printf("Decoded message: %s|%s\n", strings.Join(frameParts[:len(frameParts)-1], "|"), gbk)
		}
	}
}

//Send hello to CSocket Server to keep client live.
func SayHelloKeepAlive(s *net.UDPConn, rcvPort int) {
	for {
		_, e := s.Write([]byte("ClientHello " + strconv.Itoa(rcvPort)))
		if e != nil {
			log.Println(e)
		}
		log.Println("Say hello to:", s.RemoteAddr().String())

		time.Sleep(time.Second * 120)
	}
}

//Receive message from http to send
func SendMsg(s *net.UDPConn) {
	for {
		msg := <-HMsg.MsgChan
		log.Println("get Message from chan:", msg)
		_, e := s.Write([]byte(msg))
		if e != nil {
			log.Println(e)
		}
	}
}

//Message fields index
type MsgFieldIdx struct {
	Prefix, QQ, GroupID, DiscussID, EncodedText, OperatedQQ int
}

const PreFixMaxSize uint = 256
const PayLoadMaxSize uint = 35768

var MapMsgIdx = map[string]MsgFieldIdx{
	"PrivateMessage":      MsgFieldIdx{Prefix: 0, QQ: 1, EncodedText: 2},
	"GroupMessage":        MsgFieldIdx{Prefix: 0, GroupID: 1, QQ: 2, EncodedText: 3},
	"DiscussMessage":      MsgFieldIdx{Prefix: 0, DiscussID: 1, QQ: 2, EncodedText: 3},
	"GroupMemberDecrease": MsgFieldIdx{Prefix: 0, GroupID: 1, QQ: 2, OperatedQQ: 3},
	"GroupMemberIncrease": MsgFieldIdx{Prefix: 0, GroupID: 1, QQ: 2, OperatedQQ: 3},
}

func main() {
	f, err := os.OpenFile("Message.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)

	}
	defer f.Close()

	log.SetOutput(f)

	// read port number from ini
	cw := ConfWorker{}
	cw.Init("./app/org.dazzyd.cqsocketapi/config.ini")
	v := cw.ValueRoom("Server", "SERVER_PORT")
	port, _ := strconv.Atoi(v)
	if !(port > 1) {
		log.Println("port error", v)
		return
	}
	cqPort := port
	rcvPort := port + 1

begin:
	s, e := net.DialUDP("udp4",
		nil,
		&net.UDPAddr{
			IP:   net.IPv4(127, 0, 0, 1),
			Port: cqPort,
		})

	if e != nil {
		log.Println(e)
		goto begin
	}

	log.Println("success connect to local upd 11253 ")

	go SayHelloKeepAlive(s, rcvPort)
	httpListenAddr := ":" + strconv.Itoa(rcvPort)
	go HMsg.StartServ(httpListenAddr)
	go SendMsg(s)

	RcvMsgFromCQ(rcvPort)
}
