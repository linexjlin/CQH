package main

import (
	"testing"
)

func Test_confWorker(t *testing.T) {
	cw := ConfWorker{}
	cw.Init("/tmp/config.ini")
	v := cw.ValueRoom("Server", "SERVER_PORT")
	t.Log(v)
}
