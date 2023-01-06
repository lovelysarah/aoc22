package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
  "sort"
	"strings"
)

func main(){
  //str,err := readInput("/sample/day-1.txt")
  str,err := readInput("/input/day-1.txt")

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

func readInput(path string) (result string, err error){
  wd, err := os.Getwd()

  file, err := os.Open(wd + path);

  if err != nil {
    return "", err
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  var str string
  for scanner.Scan() {
    str += scanner.Text() + "\n"
  }

  return str, nil
}
