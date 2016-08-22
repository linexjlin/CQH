package cqencode

import (
	"encoding/base64"
	"github.com/axgle/mahonia"
	"log"
)

//convert utf8 to base64
func Utf8_gbk_base64(src string) string {
	ec := mahonia.NewEncoder("gbk")
	gbk := ec.ConvertString(src)
	b64 := base64.StdEncoding.EncodeToString([]byte(gbk))
	return b64
}

//convert base64 to utf8
func Base64_gbk_utf8(b64 string) string {
	gbk, e := base64.StdEncoding.DecodeString(b64)
	if e != nil {
		log.Println(e)
	}
	dc := mahonia.NewDecoder("gbk")
	utf8 := dc.ConvertString(string(gbk))

	return utf8
}

//convert base64 to gbk
func Base64_gbk(b64 string) string {
	gbk, e := base64.StdEncoding.DecodeString(b64)
	if e != nil {
		log.Println(e)
	}
	return string(gbk)
}

//convert utf8 to gbk
func Utf8_gbk(src string) string {
	ec := mahonia.NewEncoder("gbk")
	gbk := ec.ConvertString(src)
	return gbk
}
