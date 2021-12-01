package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2016, 5, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	base := strings.TrimSpace(input)
	password := [8]byte{}
	nonce := -1
	var c byte
	for i := range password {
		nonce, c, _ = nextHash(base, nonce)
		password[i] = c
	}

	return string(password[:])
}

// Implement Solution to Problem 2
func solve2(input string) string {
	base := strings.TrimSpace(input)
	password := [8]byte{}
	nonce := -1
	var c, p byte
	found := 0
	for found < 8 {
		nonce, p, c = nextHash(base, nonce)
		if p > 47 && p < 56 && password[p-48] == 0 {
			password[p-48] = c
			found++
		}
	}
	return string(password[:])
}

func nextHash(base string, startNonce int) (nonce int, c1, c2 byte) {
	nonce = startNonce + 1
	h := md5.New()
	for {
		fmt.Fprintf(h, "%s%d", base, nonce)
		b := h.Sum(nil)
		if b[0] == 0 && b[1] == 0 && b[2] < 16 {
			// first 5 hex digits are 0
			// the next chars are the hex decoding
			s := hex.EncodeToString(b[2:4])
			// s[0] === 0
			return nonce, s[1], s[2]
		}
		h.Reset()
		nonce++
	}
}
