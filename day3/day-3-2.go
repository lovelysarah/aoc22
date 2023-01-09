package day3

import (
	"aoc22/utils"
	"fmt"
	"strings"
)

func Part2() {
	fmt.Println("-------- Day 3 Part 2 ---------")
	//str, err := utils.ReadInput("/sample/day-3.txt")
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

	var groups [][]string
	var group []string
	for i, line := range lines {

		group = append(group, line)
		if i%3 == 2 {
			groups = append(groups, group)
			group = []string{}
		}
	}
	var result []rune

	for _, group := range groups {

		first, second, third := group[0], group[1], group[2]

		common := make(map[rune]struct{})

		for _, r1 := range first {
			if _, ok := common[r1]; ok {
				continue
			}
			for _, r2 := range second {
				if r1 != r2 {
					continue
				}
				for _, r3 := range third {
					if r2 == r3 {
						common[r1] = struct{}{}
					}
				}
			}
		}
		for k := range common {
			result = append(result, k)
		}

	}

	priorities := 0
	for _, char := range result {
		for idx, c := range alphabet {
			if char == c {
				priorities += idx + 1
			}
		}
	}

	// Add +2 because of empty lines
	fmt.Printf("Priority sum: %v\n", priorities)
}
