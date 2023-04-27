package base62

import (
	"errors"
	"strings"
)

func encode(val int64) (string, error) {
	if val < 0 {
		return "", errors.New("param error,不支持负数")
	}
	builder := &strings.Builder{}
	for val > 0 {
		num := val % 62
		val /= 62
		builder.WriteRune(rune(num))
	}
	return builder.String(), nil
}
