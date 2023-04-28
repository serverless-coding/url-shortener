package base62

import (
	"errors"
	"strings"
)

// base 62 characters
var _table = [62]rune{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

/*
0-0, ..., 9-9,
10-a, 11-b, ..., 35-z,
36-A, ..., 61-Z,
where ‘a’ stands for 10, ‘Z’ stands for 61

26+26+10
*/
func genTable() {
	id := 0
	for i := '0'; i <= '9'; i++ {
		_table[id] = i
		id++
	}
	for i := 'a'; i <= 'z'; i++ {
		_table[id] = i
		id++
	}
	for i := 'A'; i <= 'Z'; i++ {
		_table[id] = i
		id++
	}
}

func Encode(val int64) (string, error) {
	if val < 0 {
		return "", errors.New("param error,不支持负数")
	}
	builder := &strings.Builder{}
	for val > 0 {
		num := val % 62
		val /= 62
		builder.WriteRune(_table[num])
	}
	return reverse(builder.String()), nil
}

func reverse(str string) string {
	if str == "" {
		return str
	}
	size := len(str)
	if size == 1 {
		return str
	}
	bs := []byte(str)
	for i, j := 0, size-1; i <= j; {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	return string(bs)
}
