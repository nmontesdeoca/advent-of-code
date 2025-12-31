package main

import (
	"log"
	"os"
	"strings"
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	for b.Loop() {
		_ = part1(lines)
	}
}

func BenchmarkPart2(b *testing.B) {
	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	for b.Loop() {
		_ = part2(lines)
	}
}
