package main

import (
	"github.com/go-ini/ini"
	"log"
)

type ConfWorker struct {
	fileName string
	cfg      *ini.File
}

func (cw *ConfWorker) Init(fileName string) {
	if fileName == "" {
		fileName = "./app/org.dazzyd.cqsocketapi/config.ini"
	}
	cw.fileName = fileName
	var err error
	cw.cfg, err = ini.Load(cw.fileName)
	if err != nil {
		log.Println(err)
	}
}

func (cw *ConfWorker) ValueRoom(section, key string) string {
	sec, e1 := cw.cfg.GetSection(section)
	if e1 != nil {
		log.Println(e1)
	}
	k, e2 := sec.GetKey(key)
	if e2 != nil {
		log.Println(e2)
	}
	v := k.Value()
	return v
}
