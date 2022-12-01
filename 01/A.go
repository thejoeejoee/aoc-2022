package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("01/in.txt")
	if err != nil {
		return
	}

	maxSum := 0
	currentSum := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if currentSum > maxSum {
				maxSum = currentSum
			}
			currentSum = 0
		} else {
			value, err := strconv.Atoi(line)
			if err != nil {
				return
			}
			currentSum += value
		}

	}
	println(maxSum)

}
