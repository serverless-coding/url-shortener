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
}

func TestEncode(t *testing.T) {
	for i, v := range _testCase {
		code, err := encode(v.in)
		if err != nil || code != v.out {
			fmt.Printf("[%d].in:%d,expect:%s,got:%s", i, v.in, v.out, code)
			t.Error(err)
		}
	}
}
