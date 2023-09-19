package base62

import (
	"fmt"
	"testing"

	"github.com/bwmarrin/snowflake"
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
	{in: 11762981932034, out: "3l5OVq8O"},
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

	fmt.Println(Encode(11_7629_8193_2034))
	fmt.Println(Encode(99_7629_8193_2034))
	fmt.Println(Encode(199762981932034))
	fmt.Println(Encode(1199762981932034))
	fmt.Println(Encode(11199762981932034))
	fmt.Println(Encode(11_1199_7629_8193_2034))
	// 3l5OVq8O
	// sknLblku
	// UIWeKluW
	// 5uGv0onbs
	// Pii6GQE2C
	// 8dimbxpmCe
	for i, v := range _testCase {
		code, err := Encode(v.in)
		if err != nil || code != v.out {
			fmt.Printf("[%d].in:%d,expect:%s,got:%s", i, v.in, v.out, code)
			t.Error(err)
		}

		val, err := Decode(code)
		if err != nil || val != v.in {
			t.Errorf("decode,not eq,got:%d,wangt:%d", val, v.in)
		}
	}
}

func TestSnowflakeId(t *testing.T) {
	node, err := snowflake.NewNode(0)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 100; i++ {
		id := node.Generate().Int64()
		uId, err := Encode(id)
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("[%d]=>%s\n", id, uId)
	}
}
