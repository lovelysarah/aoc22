package day2

import (
	"aoc22/utils"
	"fmt"
	"strings"
)

type Outcomes struct {
	Z string
	Y string
	X string
}

type Pick struct {
	Outcome Outcomes
}

type RockPaperScissor struct {
	Rock    Pick
	Paper   Pick
	Scissor Pick
}

func Run() {

	fmt.Println("-------- Day 2 ---------")
	str, err := utils.ReadInput("/input/day-2.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(str, "\n")

	score := 0

	var rock, paper, scissor string = "A", "B", "C"
	var lose, draw, win string = "X", "Y", "Z"

	for _, line := range lines {

		if len(line) == 0 {
			continue
		}

		play := line[2:3]
		against := line[0:1]

		//Part 2
		mapping := RockPaperScissor{
			Rock: Pick{
				Outcome: Outcomes{
					Z: paper,
					Y: rock,
					X: scissor,
				},
			},
			Paper: Pick{
				Outcome: Outcomes{
					Z: scissor,
					Y: paper,
					X: rock,
				},
			},
			Scissor: Pick{
				Outcome: Outcomes{
					Z: rock,
					Y: scissor,
					X: paper,
				},
			},
		}

		var pick Pick

		switch against {
		case "A":
			pick = mapping.Rock
		case "B":
			pick = mapping.Paper
		case "C":
			pick = mapping.Scissor
		}

		// Part 2
		switch play {
		case win:
			play = pick.Outcome.Z
		case draw:
			play = pick.Outcome.Y
		case lose:
			play = pick.Outcome.X
		}

		// Part 1
		switch play {
		case rock: // Rock
			score += calculatePoints(scissor, rock, 1, against)
		case paper: // Paper
			score += calculatePoints(rock, paper, 2, against)
		case scissor: // Scissor
			score += calculatePoints(paper, scissor, 3, against)
		}

	}
	fmt.Printf("Score: %v\n", score)
}

func calculatePoints(win string, draw string, bonus int, against string) (points int) {
	point := 0
	if against == win {
		point += bonus + 6
	} else if against == draw {
		point += bonus + 3
	} else {
		point += bonus + 0
	}
	return point
}
