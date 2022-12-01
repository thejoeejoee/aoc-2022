package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("01/in.txt")
	if err != nil {
		return
	}

	currentSum := 0

	scanner := bufio.NewScanner(file)
	var values []int

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			values = append(values, currentSum)
			currentSum = 0
		} else {
			value, err := strconv.Atoi(line)
			if err != nil {
				return
			}
			currentSum += value
		}

	}

	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	println(values[0] + values[1] + values[2])
}
