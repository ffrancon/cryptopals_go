package set1

import (
	"ffrancon/cryptopals-go/pkg"
	"testing"
)

func TestChallenge1(t *testing.T) {
	hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	exp := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	res, _ := pkg.HexStrToBase64(hex)
	if string(res) != exp {
		t.Errorf("expected %s, got %s", exp, res)
	}
}

func TestChallenge2(t *testing.T) {
	buf1 := "1c0111001f010100061a024b53535009181c"
	buf2 := "686974207468652062756c6c277320657965"
	exp := "746865206b696420646f6e277420706c6179"
	res := pkg.XorBuffers(buf1, buf2)
	if res != exp {
		t.Errorf("expected %s, got %s", exp, res)
	}
}

func TestChallenge3(t *testing.T) {
	str := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	m := "Cooking MC's like a pound of bacon"
	key := byte(88)
	res := pkg.DecryptXorSingleByte(str)
	if string(res.Decrypted[:]) != m && res.Key != key {
		t.Errorf("expected %s & %d, got %s & %d", m, key, res.Decrypted, res.Key)
	}
}
