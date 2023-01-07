package day2

import (
	"aoc22/utils"
	"fmt"
	"strings"
)

func Run() {

	fmt.Println("-------- Day 2 ---------")
	// str, err := utils.ReadInput("/sample/day-2.txt")
	str, err := utils.ReadInput("/input/day-2.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(str, "\n")

	score := 0

	var rock, paper, scissor string = "A", "B", "C"
	for idx, line := range lines {

		if len(line) == 0 {
			return
		}

		play := line[2:3]
		against := line[0:1]

		switch play {
		case "X": // Rock
			score += calculatePoints(scissor, rock, 1, against)
		case "Y": // Paper
			score += calculatePoints(rock, paper, 2, against)
		case "Z": // Scissor
			score += calculatePoints(paper, scissor, 3, against)
		}

		fmt.Printf("%v: %v | score: %v\n", idx, line, score)
	}
}

func calculatePoints(win string, draw string, bonus int, against string) (points int) {
	point := 0
	if against == win {
		fmt.Println("win")
		point += bonus + 6
	} else if against == draw {
		fmt.Println("draw")
		point += bonus + 3
	} else {
		fmt.Println("Lost")
		point += bonus + 0
	}
	return point
}
