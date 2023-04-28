package base62

import (
	"fmt"
	"testing"
)

var _testCase = []struct {
	in  int64
	out string
}{
	{in: 1, out: "1"},
	{in: 11157, out: "2TX"},
	{in: 62, out: "10"},
	{in: 61, out: "Z"},
	{in: 63, out: "11"},
	{in: 11158, out: "2TY"},
	{in: 11159, out: "2TZ"},
}

func TestGenTable(t *testing.T) {
	genTable()

	for i, v := range _table {
		fmt.Printf("'%c',", v)
		if i == 30 {
			fmt.Println()
		}
	}
}

func TestEncode(t *testing.T) {
	for i, v := range _testCase {
		code, err := Encode(v.in)
		if err != nil || code != v.out {
			fmt.Printf("[%d].in:%d,expect:%s,got:%s", i, v.in, v.out, code)
			t.Error(err)
		}
	}
}
