package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	fmt.Println("part 1", part1(lines))
	fmt.Println("part 2", part2(lines))
}

func part1(lines []string) *big.Int {
	return sumSelected(lines, 2)
}

func part2(lines []string) *big.Int {
	return sumSelected(lines, 12)
}

// maxSubsequence returns the lexicographically largest subsequence of length k
// from the digit string s (keeps original order). Assumes s contains only '0'..'9'.
// If k >= len(s) it returns s.
func maxSubsequence(s string, k int) string {
	n := len(s)
	if k >= n {
		return s
	}
	stack := make([]byte, 0, k)
	for i := 0; i < n; i++ {
		ch := s[i]
		rem := n - i // remaining characters including current
		// pop while we can replace a smaller previous digit with a larger current one
		for len(stack) > 0 && stack[len(stack)-1] < ch && (len(stack)-1+rem) >= k {
			stack = stack[:len(stack)-1]
		}
		if len(stack) < k {
			stack = append(stack, ch)
		}
	}
	// safety: trim if somehow longer
	if len(stack) > k {
		stack = stack[:k]
	}
	return string(stack)
}

// sumSelected takes a slice of digit-strings and picks the max subsequence
// of length k from each, parsing them as big.Int and summing. Returns the sum.
func sumSelected(lines []string, k int) *big.Int {
	total := big.NewInt(0)
	for _, line := range lines {
		selected := maxSubsequence(line, k)
		v := new(big.Int)
		if _, ok := v.SetString(selected, 10); !ok {
			// Shouldn't happen for digit-only strings
			panic("failed to parse big integer from selected digits: " + selected)
		}
		total.Add(total, v)
	}
	return total
}
