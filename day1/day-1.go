package day1

import (
	"aoc22/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Run() {
	str, err := utils.ReadInput("/input/day-1.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(str, "\n")
	groups := [][]int{}

	var group []int

	for _, line := range lines {
		if line == "" {
			groups = append(groups, group)
			group = []int{}
		} else {
			num, _ := strconv.Atoi(line)

			group = append(group, num)
		}
	}

	var sums []int
	for _, group := range groups {
		sum := 0
		for _, element := range group {
			sum += element
		}
		sums = append(sums, sum)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	// Part 1
	fmt.Println(sums[0])

	// Part 2
	biggests := 0
	for _, x := range sums[:3] {
		biggests += x
	}

	fmt.Println(biggests)
}
