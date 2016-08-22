package HMsg

import (
	"cqencode"
	"log"
	"net/http"
	"time"
)

var MsgChan = make(chan string, 1000)

func msgSend(w http.ResponseWriter, r *http.Request) {
	msg := ""
	switch r.Method {
	case "GET":
		log.Println("Get")

	case "POST":
		r.ParseForm()
		log.Println("Post")
	default:
		log.Println("others request method, not support!")
		log.Println("the method is:", r.Method)
		http.NotFound(w, r)
		return
	}

	Prefix := r.FormValue("Prefix")
	switch Prefix {
	case "PrivateMessage":
		QQ := r.FormValue("QQ")
		Text := r.FormValue("Text")
		EncodedText := cqencode.Utf8_gbk_base64(Text)
		msg = Prefix + " " + QQ + " " + EncodedText
	case "GroupMessage":
		GroupID := r.FormValue("GroupID")
		Text := r.FormValue("Text")
		EncodedText := cqencode.Utf8_gbk_base64(Text)
		msg = Prefix + " " + GroupID + " " + EncodedText
	case "DiscussMessage":
		DiscussID := r.FormValue("DiscussID")
		Text := r.FormValue("Text")
		EncodedText := cqencode.Utf8_gbk_base64(Text)
		msg = Prefix + " " + DiscussID + " " + EncodedText
	default:
		log.Println("Prefix wrong!")
		w.Write([]byte("U prefix wrong!"))
		return
	}

	MsgChan <- msg
	w.Write([]byte("OK\n"))
}

func StartServ() {
	http.HandleFunc("/send", msgSend)

	srv := &http.Server{
		Addr:         ":11235",
		ReadTimeout:  time.Second * 90,
		WriteTimeout: time.Second * 90,
	}
	srv.ListenAndServe()
}
