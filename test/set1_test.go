package set1

import (
	"ffrancon/cryptopals-go/set1"
	"testing"
)

func TestChal1(t *testing.T) {
	hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	exp := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	res := set1.Chal1(hex)
	if string(res) != exp {
		t.Errorf("expected %s, got %s", exp, res)
	}
}

func TestChal2(t *testing.T) {
	buf1 := "1c0111001f010100061a024b53535009181c"
	buf2 := "686974207468652062756c6c277320657965"
	exp := "746865206b696420646f6e277420706c6179"
	res := set1.Chal2(buf1, buf2)
	if res != exp {
		t.Errorf("expected %s, got %s", exp, res)
	}
}
