package main

import (
	"bufio"
	"os"
)

const (
	ROCK_0     = "A "
	PAPER_0    = "B "
	SCISSORS_0 = "C "

	ROCK_1     = "X"
	PAPER_1    = "Y"
	SCISSORS_1 = "Z"
)

func main() {
	file, err := os.Open("02/in.txt")
	if err != nil {
		return
	}

	m := make(map[string]int)
	m[ROCK_0+ROCK_1] = 1 + 3
	m[ROCK_0+PAPER_1] = 2 + 6
	m[ROCK_0+SCISSORS_1] = 3

	m[PAPER_0+ROCK_1] = 1
	m[PAPER_0+PAPER_1] = 2 + 3
	m[PAPER_0+SCISSORS_1] = 3 + 6

	m[SCISSORS_0+ROCK_1] = 1 + 6
	m[SCISSORS_0+PAPER_1] = 2
	m[SCISSORS_0+SCISSORS_1] = 3 + 3

	scanner := bufio.NewScanner(file)

	var sum = 0

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			sum += m[line]
		}

	}
	println(sum)

}
