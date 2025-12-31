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

	lines := strings.Split(string(data), ",")

	fmt.Println("part 1", part1(lines))
	fmt.Println("part 2", part2(lines))
}

func part1(lines []string) int {
	var sumInvalidIds int

	for _, line := range lines {
		lineParts := strings.Split(line, "-")
		startStr, endStr := lineParts[0], lineParts[1]
		start, _ := strconv.Atoi(startStr)
		end, _ := strconv.Atoi(endStr)

		for i := start; i <= end; i++ {
			if !isValidIdPart1(strconv.Itoa(i)) {
				sumInvalidIds += i
			}
		}
	}

	return sumInvalidIds
}

func isValidIdPart1(id string) bool {
	middle := len(id) / 2

	return id[:middle] != id[middle:]
}

func part2(lines []string) int {
	var sumInvalidIds int

	for _, line := range lines {
		lineParts := strings.Split(line, "-")
		startStr, endStr := lineParts[0], lineParts[1]
		start, _ := strconv.Atoi(startStr)
		end, _ := strconv.Atoi(endStr)

		for i := start; i <= end; i++ {
			idStr := strconv.Itoa(i)
			valid := isValidIdPart2(idStr)

			if !valid {
				sumInvalidIds += i
			}
		}
	}

	return sumInvalidIds
}

func isValidIdPart2(id string) bool {
	idLength := len(id)

	for i := 1; i <= idLength/2; i++ {
		if idLength%i == 0 {
			part := id[:i]

			if strings.Repeat(part, idLength/i) == id {
				return false
			}
		}
	}

	return true
}
