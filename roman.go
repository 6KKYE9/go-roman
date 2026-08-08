package main

import (
	"errors"
	"strings"
)

var ErrRoman = errors.New("不是合法的罗马数字")

// 从大到小排，处理 IV/IX 这种减法组合
var table = []struct {
	sym string
	val int
}{
	{"M", 1000}, {"CM", 900}, {"D", 500}, {"CD", 400},
	{"C", 100}, {"XC", 90}, {"L", 50}, {"XL", 40},
	{"X", 10}, {"IX", 9}, {"V", 5}, {"IV", 4}, {"I", 1},
}

// ToRoman 把 1~3999 的整数转成罗马数字，超出范围返回空串
func ToRoman(n int) string {
	if n <= 0 || n > 3999 {
		return ""
	}
	var sb strings.Builder
	for _, v := range table {
		for n >= v.val {
			n -= v.val
			sb.WriteString(v.sym)
		}
	}
	return sb.String()
}

// FromRoman 把罗马数字转成整数，字符不合法时报错
func FromRoman(s string) (int, error) {
	s = strings.ToUpper(strings.TrimSpace(s))
	if s == "" {
		return 0, ErrRoman
	}
	total, prev := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		var cur int
		switch s[i] {
		case 'I':
			cur = 1
		case 'V':
			cur = 5
		case 'X':
			cur = 10
		case 'L':
			cur = 50
		case 'C':
			cur = 100
		case 'D':
			cur = 500
		case 'M':
			cur = 1000
		default:
			return 0, ErrRoman
		}
		// 小的在大的左边要减（如 IV=4），否则加
		if cur < prev {
			total -= cur
		} else {
			total += cur
		}
		prev = cur
	}
	return total, nil
}
