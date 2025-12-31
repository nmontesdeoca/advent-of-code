package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

func part1(lines []string) int {
	currentPosition := 50
	countPositionZero := 0

	for _, v := range lines {
		right := v[0] == 'R'
		count, err := strconv.Atoi(v[1:])

		if err != nil {
			log.Fatal(err)
		}

		if right {
			currentPosition += count
		} else {
			currentPosition -= count
		}

		if currentPosition < 0 {
			for currentPosition < 0 {
				currentPosition = 100 + currentPosition
			}
		} else if currentPosition > 99 {
			for currentPosition > 99 {
				currentPosition = currentPosition - 100
			}
		}

		if currentPosition == 0 {
			countPositionZero++
		}
	}

	return countPositionZero
}

func part2(lines []string) int {
	currentPosition := 50
	countPositionZero := 0

	// fmt.Println("The dial starts by pointing at 50.")

	for _, v := range lines {
		right := v[0] == 'R'
		count, err := strconv.Atoi(v[1:])

		if err != nil {
			log.Fatal(err)
		}

		if right {
			for i := 1; i <= count; i++ {
				if currentPosition == 99 {
					currentPosition = 0
					countPositionZero++
				} else {
					currentPosition++
				}
			}
		} else {
			for i := 1; i <= count; i++ {
				if currentPosition == 0 {
					currentPosition = 99
					if i != 1 {
						countPositionZero++
					}
				} else {
					currentPosition--
				}
			}
			if currentPosition == 0 {
				countPositionZero++
			}
		}

		// fmt.Printf("The dial is rotated %v to point at %v; during this rotation, it points at 0 %v\n", v, currentPosition, countPositionZero)
	}

	return countPositionZero
}
