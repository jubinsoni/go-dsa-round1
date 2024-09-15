package main

import (
  "fmt"
  "slices"
)

func main() {
  s := []string{"abc", "def", "ghi"}
  var x = "def"
  var y = "lmn"

  fmt.Println(slices.Contains(s, x))
  fmt.Println(slices.Contains(s, y))

  x1 := []int{}
  x1 = append(x1, 5)
  fmt.Println(x1)
}
