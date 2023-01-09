package day3

import (
	"aoc22/utils"
	"fmt"
	"strings"
)

func Part1() {
	fmt.Println("-------- Day 3 Part 1 ---------")
	str, err := utils.ReadInput("/input/day-3.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(str, "\n")

	alphabet := []rune{}
	for i := 'a'; i <= 'z'; i++ {
		alphabet = append(alphabet, i)
	}
	for i := 'A'; i <= 'Z'; i++ {
		alphabet = append(alphabet, i)
	}

	var matches []rune
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		half := len(line) / 2

		first, second := line[0:half], line[half:]

		flag := false
		for _, r1 := range first {
			for _, r2 := range second {
				if r1 == r2 && !flag {
					matches = append(matches, r1)
					flag = true
				}
			}
		}

		priorities := 0
		for _, char := range matches {
			for idx, c := range alphabet {
				if char == c {
					priorities += idx + 1
				}
			}
		}

		// Add +2 because of empty lines
		if len(lines) == i+2 {
			fmt.Printf("Priority sum: %v\n", priorities)
		}
	}
}
