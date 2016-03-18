package main

import (
  "fmt"
  "strings"
  )

func main() {
  f := "45,36,83"
  a := strings.Split(f,",")
  fmt.Println(a)
  for i := 0; i<len(a); {
    fmt.Println(a[i])
    i = i + 1
  }
}
