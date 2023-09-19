package base62

import (
	"errors"
	"math"
	"strings"
	"sync"
)

const (
	_base   = 62
	_maxLen = 15
)

// base 62 characters 62^6=>560亿 11位
var (
	_table = [_base]rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

	power = [_maxLen]int64{
		int64(math.Pow(62, 0)),
		int64(math.Pow(62, 1)),
		int64(math.Pow(62, 2)),
		int64(math.Pow(62, 3)),
		int64(math.Pow(62, 4)),
		int64(math.Pow(62, 5)),
		int64(math.Pow(62, 6)),
		int64(math.Pow(62, 7)),
		int64(math.Pow(62, 8)),
		int64(math.Pow(62, 9)),
		int64(math.Pow(62, 10)),
		int64(math.Pow(62, 11)),
		int64(math.Pow(62, 12)),
		int64(math.Pow(62, 13)),
		int64(math.Pow(62, 14)),
	}
	_mTable map[rune]int = nil
	_mu                  = sync.Mutex{}
)

func Decode(uId string) (int64, error) {
	if _mTable == nil {
		_mu.Lock()
		defer _mu.Unlock()
		if _mTable == nil {
			_mTable = make(map[rune]int, 62)
			for i, v := range _table {
				_mTable[v] = i
				// fmt.Printf("%c:%d\n", v, i)
			}
		}
	}

	res, n := int64(0), len(uId)
	for i, v := range uId {
		res += int64(_mTable[v]) * power[n-1-i]
	}
	return res, nil
}

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
