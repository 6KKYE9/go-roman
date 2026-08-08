package main

import "testing"

func TestToRoman(t *testing.T) {
	cases := map[int]string{
		1: "I", 4: "IV", 9: "IX", 40: "XL", 1994: "MCMXCIV", 3999: "MMMCMXCIX", 0: "",
	}
	for n, want := range cases {
		if got := ToRoman(n); got != want {
			t.Errorf("ToRoman(%d)=%q 想要 %q", n, got, want)
		}
	}
}

func TestFromRoman(t *testing.T) {
	cases := map[string]int{
		"I": 1, "IV": 4, "IX": 9, "XL": 40, "MCMXCIV": 1994, "MMMCMXCIX": 3999,
	}
	for s, want := range cases {
		got, err := FromRoman(s)
		if err != nil {
			t.Fatalf("FromRoman(%q) 出错: %v", s, err)
		}
		if got != want {
			t.Errorf("FromRoman(%q)=%d 想要 %d", s, got, want)
		}
	}
}

func TestFromRomanBad(t *testing.T) {
	// 宽松解析：只要字符合法就算，所以只测含非法字符的情况
	for _, s := range []string{"", "ABC", "MXQ", "123"} {
		if _, err := FromRoman(s); err == nil {
			t.Errorf("FromRoman(%q) 该报错却没报", s)
		}
	}
}

func TestRoundTrip(t *testing.T) {
	for n := 1; n <= 3999; n++ {
		r := ToRoman(n)
		back, err := FromRoman(r)
		if err != nil || back != n {
			t.Errorf("往返不一致: %d -> %q -> %d", n, r, back)
		}
	}
}
