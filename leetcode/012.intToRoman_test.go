package leetcode

import "testing"

func Test_intToRoman(t *testing.T) {
	r := intToRoman(3)
	if r != "III" {
		t.Error("III", r)
	}
	r = intToRoman(4)
	if r != "IV" {
		t.Error("IV", r)
	}
	r = intToRoman(9)
	if r != "IX" {
		t.Error("IX", r)
	}
	r = intToRoman(58)
	if r != "LVIII" {
		t.Error("LVIII", r)
	}
	r = intToRoman(1994)
	if r != "MCMXCIV" {
		t.Error("MCMXCIV", r)
	}
}
