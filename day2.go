
package main

import (
  "fmt"
  // "reflect"
)

func abs(x int) int {
  if x > 0 {
    return x
  }

  return -x
}

func main(){
  var s string = "hello"
  fmt.Println(s)

  for pos, ch := range(s) {
    fmt.Println(pos, ch)
  }

  fmt.Println("-------")

  var res int
  for i:=1; i<len(s); i++ {
    fmt.Println(i, abs(int(s[i-1])-int(s[i])))
    res = res + abs(int(s[i-1])-int(s[i]))
  }

  fmt.Println(res)
}
