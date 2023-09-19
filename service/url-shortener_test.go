package service

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/ini.v1"
)

func init() {
	cfg, err := ini.Load("../.env")
	if err != nil {
		panic(err)
	}
	for _, v := range cfg.Section("").Keys() {
		os.Setenv(v.Name(), v.String())
	}
}

func TestBase62(t *testing.T) {
	s := NewUrlShortener()

	code, err := s.Short("https://www.aliyun.com")
	if err != nil {
		t.Errorf("%s,%v", code, err)
		return
	}
	url, err := s.UrlFromShort(code)
	fmt.Println(url, err)
}
