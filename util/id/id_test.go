package id

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/ini.v1"
)

func init() {
	cfg, err := ini.Load("../../.env")
	if err != nil {
		panic(err)
	}
	for _, v := range cfg.Section("").Keys() {
		os.Setenv(v.Name(), v.String())
	}
}

func TestId(t *testing.T) {
	fmt.Printf("sequenceMask=%d,binary=%b,next=%b\n",
		sequenceMax, sequenceMax, sequenceMax+1)

	s, err := NewSnowflake(0, 1, true)
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(s.Id())
	}

	//
	s, err = NewSnowflake(0, 1, false)
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(s.Id())
	}
}
