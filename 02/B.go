package main

import (
	"bufio"
	"os"
)

const (
	ROCK_0     = "A "
	PAPER_0    = "B "
	SCISSORS_0 = "C "

	LOSE = "X"
	DRAW = "Y"
	WIN  = "Z"

	ROCK     = 1
	PAPER    = 2
	SCISSORS = 3
)

func main() {
	// x lose
	// y draw
	// z win

	file, err := os.Open("02/in.txt")
	if err != nil {
		return
	}

	m := make(map[string]int)
	m[ROCK_0+LOSE] = SCISSORS
	m[ROCK_0+DRAW] = ROCK + 3
	m[ROCK_0+WIN] = PAPER + 6

	m[PAPER_0+LOSE] = ROCK
	m[PAPER_0+DRAW] = PAPER + 3
	m[PAPER_0+WIN] = SCISSORS + 6

	m[SCISSORS_0+LOSE] = PAPER
	m[SCISSORS_0+DRAW] = SCISSORS + 3
	m[SCISSORS_0+WIN] = ROCK + 6

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
