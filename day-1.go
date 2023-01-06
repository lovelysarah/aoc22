package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
  //str,err := readInput("/sample/day-1.txt")
  str,err := readInput("/input/day-1.txt")

  if err != nil {
    return
  }

  fmt.Println(str)

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


  for idx, group := range groups {
    sum := 0
    for _, element := range group {
      sum += element
      groups[idx] = []int{sum}
    }
  }

  largest := 0
  for _, group := range groups {
    if group[0] > largest {
      largest = group[0]
    }
  }

  //Answer
  fmt.Println(largest)

}

func readInput(path string) (result string, err error){
  wd, err := os.Getwd()

  file, err := os.Open(wd + path);

  if err != nil {
  fmt.Println(err)
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
