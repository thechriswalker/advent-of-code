package main

import (
	"crypto/md5"
	"fmt"
	"hash"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2016, 14, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	base := strings.TrimSpace(input)
	h := md5.New()
	nonce := 0
	cache := map[int]string{}
	found := 0
	var hash string
	var ok bool
	var a byte
	for found < 64 {
		nonce++
		if hash, ok = cache[nonce]; !ok {
			h.Reset()
			fmt.Fprintf(h, "%s%d", base, nonce)
			hash = fmt.Sprintf("%x", h.Sum(nil))
			cache[nonce] = hash
		}
		if a, ok = hasTriple(hash); ok {
			//	fmt.Printf("found triple %c in hash %s at %d\n", a, hash, nonce)
			for i := 1; i < 1001; i++ {
				if hash, ok = cache[nonce+i]; !ok {
					h.Reset()
					fmt.Fprintf(h, "%s%d", base, nonce+i)
					hash = fmt.Sprintf("%x", h.Sum(nil))
					cache[nonce+i] = hash
				}
				if hasQuint(hash, a) {
					//	fmt.Printf("found quint %c in hash %s at %d\n", a, hash, nonce+i)
					found++
					break
				}
			}
		}
		delete(cache, nonce)
	}
	return fmt.Sprintf("%d", nonce)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	base := strings.TrimSpace(input)
	h := md5.New()
	nonce := 0
	cache := map[int]string{}
	found := 0
	var hash string
	var ok bool
	var a byte
	for found < 64 {
		nonce++
		if hash, ok = cache[nonce]; !ok {
			hash = Hash2016(h, base, nonce)
			cache[nonce] = hash
		}
		if a, ok = hasTriple(hash); ok {
			//		fmt.Printf("found triple %c in hash %s at %d\n", a, hash, nonce)
			for i := 1; i < 1001; i++ {
				if hash, ok = cache[nonce+i]; !ok {
					hash = Hash2016(h, base, nonce+i)
					cache[nonce+i] = hash
				}
				if hasQuint(hash, a) {
					//						fmt.Printf("found quint %c in hash %s at %d\n", a, hash, nonce+i)
					found++
					break
				}
			}
		}
		delete(cache, nonce)
	}
	return fmt.Sprintf("%d", nonce)
}

func Hash2016(h hash.Hash, base string, nonce int) string {
	h.Reset()
	fmt.Fprintf(h, "%s%d", base, nonce)
	b := h.Sum(nil)
	for i := 0; i < 2016; i++ {
		h.Reset()
		fmt.Fprintf(h, "%x", b)
		b = h.Sum(nil)
	}
	return fmt.Sprintf("%x", b)
}

func hasTriple(hash string) (a byte, ok bool) {
	for i := 0; i < len(hash)-2; i++ {
		if hash[i] == hash[i+1] && hash[i] == hash[i+2] {
			return hash[i], true
		}
	}
	return 'x', false
}

func hasQuint(hash string, a byte) bool {
	for i := 0; i < len(hash)-4; i++ {
		if hash[i] == a && hash[i+1] == a && hash[i+2] == a && hash[i+3] == a && hash[i+4] == a {
			return true
		}
	}
	return false
}
