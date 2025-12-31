package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input_test.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	fmt.Println("part 1", part1(lines))
	fmt.Println("part 2", part2(lines))
}

func part1(lines []string) int {
	totalJoltage := 0

	for _, line := range lines {
		var lineJoltage int
		// n and m will be my 2 numbers forming number nm
		lineLength := len(line)
		n := line[lineLength-2]
		m := line[lineLength-1]

		for i := lineLength - 3; i >= 0; i-- {
			if line[i] >= n {
				if n >= m {
					m = n
				}
				n = line[i]
			}
		}

		lineJoltage, _ = strconv.Atoi(string([]byte{n, m}))
		totalJoltage += lineJoltage
	}

	return totalJoltage
}

func part2(lines []string) int {
	totalJoltage := 0

	for _, line := range lines {
		var i int
		var v byte
		lineJoltageBytes := make([]byte, 0, 12)

		i, v = getHighestJoltage(line[0 : len(line)-11])
		lineJoltageBytes = append(lineJoltageBytes, v)

		i, v = getHighestJoltage(line[i+1 : len(line)-10])
		lineJoltageBytes = append(lineJoltageBytes, v)

		i, v = getHighestJoltage(line[i+1 : len(line)-9])
		lineJoltageBytes = append(lineJoltageBytes, v)

		i, v = getHighestJoltage(line[i+1 : len(line)-8])
		lineJoltageBytes = append(lineJoltageBytes, v)

		i, v = getHighestJoltage(line[i+1 : len(line)-7])
		lineJoltageBytes = append(lineJoltageBytes, v)

		i, v = getHighestJoltage(line[i+1 : len(line)-6])
		lineJoltageBytes = append(lineJoltageBytes, v)

		i, v = getHighestJoltage(line[i+1 : len(line)-5])
		lineJoltageBytes = append(lineJoltageBytes, v)

		i, v = getHighestJoltage(line[i+1 : len(line)-4])
		lineJoltageBytes = append(lineJoltageBytes, v)

		i, v = getHighestJoltage(line[i+1 : len(line)-3])
		lineJoltageBytes = append(lineJoltageBytes, v)

		i, v = getHighestJoltage(line[i+1 : len(line)-2])
		lineJoltageBytes = append(lineJoltageBytes, v)

		i, v = getHighestJoltage(line[i+1 : len(line)-1])
		lineJoltageBytes = append(lineJoltageBytes, v)

		i, v = getHighestJoltage(line[i+1:])
		lineJoltageBytes = append(lineJoltageBytes, v)
		fmt.Println(string(lineJoltageBytes))
		lineJoltage, _ := strconv.Atoi(string(lineJoltageBytes))
		totalJoltage += lineJoltage
	}

	return totalJoltage
}

func getHighestJoltage(line string) (int, byte) {
	var max byte
	var index int

	for i, v := range line {
		if byte(v) > max {
			max = byte(v)
			index = i

		}
	}

	fmt.Println("line", line, "max", string(max), "index", index)

	return index, max
}
