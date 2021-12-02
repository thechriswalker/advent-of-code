package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"math"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 4, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	key := strings.TrimSpace(input)
	return solve([]byte(key), []byte{0x00, 0x00, 0x0f})
}
func solve(key, min []byte) string {
	// need to find the lowest int x that when converted
	// to a string and concatenated to the key, then MD5 hashed
	// produces an output with at least 5 zeroes at the start in hex.
	// e.g. 00 00 01 dbbfa....
	// now MD5 hashes are 128bits, and the first 3 bytes must be less than 0x00000F
	// so we can bytes.Compare and not keep re-hexing
	i := uint64(0)
	h := md5.New()
	sum := make([]byte, 0, md5.Size)
	l := len(min)
	for {
		i++
		h.Reset()
		h.Write(key)
		fmt.Fprintf(h, "%d", i)

		sum = h.Sum(sum[0:0])
		// if i == 609043 {
		// 	fmt.Printf("i = %d, toHash = %s%d, sum = %x, sum[:%d] = %v, cmp = %d\n", i, key, i, sum, sum[:l], l, bytes.Compare(min, sum[:l]))
		// }
		if bytes.Compare(min, sum[:l]) > -1 {
			return fmt.Sprintf("%d", i)
		}
		//time.Sleep(200 * time.Millisecond)
		if i == math.MaxUint64 {
			panic("NOT FOUND!")
		}
	}
}

// Implement Solution to Problem 2
func solve2(input string) string {
	key := strings.TrimSpace(input)
	return solve([]byte(key), []byte{0x00, 0x00, 0x00})
}
