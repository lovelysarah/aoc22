package utils

import (
  "os"
  "bufio"
)

func ReadInput(path string) (result string, err error){
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
