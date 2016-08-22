package HMsg

import (
	//"fmt"
	"testing"
)

func Test_utf8_gbk_base64(t *testing.T) {
	src := "test 测试"
	b64 := utf8_gbk_base64(src)
	tsrc := base64_gbk_utf8(b64)
	if tsrc != src {
		t.Error("encode or decode error!")
	}
}

func TestStartServ(t *testing.T) {
	//b64 := utf8_gbk_base64("test 测试")
	//	StartServ()
}
