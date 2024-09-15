package main

import (
  "fmt"
  "reflect"
  "maps" // only for using maps.Equal
)

// declaring constants
const ENV string = "DEV"

func main(){
  // declaring array and looping over it
  var strz = []string{"abc", "def", "ghi", "jkl"}

  fmt.Println(strz)

  for ind, str := range(strz){
    fmt.Println(ind, str)
  }

  // variables declaration and type
  var pi float64 = 3.1467
  var days_in_week int64 = 7
  var curr_value int = 2

  fmt.Println(pi, reflect.TypeOf(pi))
  fmt.Println(days_in_week, reflect.TypeOf(days_in_week))
  fmt.Println(curr_value, reflect.TypeOf(curr_value))

  // string formatting
  var env = fmt.Sprintf("Current Environment is %s", ENV)
  fmt.Println(env)

  // declare without initialization 
  // by default 0 value is assined
  var initz int64
  var floatz float64
  fmt.Println(initz, floatz)

  fmt.Println("--------------")

  // declare without initialization 
  // by default [0, 0, 0, 0, 0] is assined
  var a [5] int
  fmt.Println(a)

  b := [5]int{10, 20, 30, 40, 50}
  fmt.Println(b)

  for i:=0; i < len(b); i++ {
    fmt.Println(b[i])
  }

  fmt.Println("--------------")

  mp := make(map[string]int)
  mp["k1"] = 7
  mp["k2"] = 14
  mp["k3"] = 20
  fmt.Println("map:", mp)

  delete(mp, "k3")
  fmt.Println("map:", mp)

  val1, exists1 := mp["k2"]
  fmt.Println(val1, exists1)

  val2, exists2 := mp["k3"]
  fmt.Println(val2, exists2)

  n := map[string]int{"foo": 1, "bar": 2}
  fmt.Println("map:", n)

  n2 := map[string]int{"foo": 1, "bar": 2}
  if maps.Equal(n, n2) {
    fmt.Println("n == n2")
  }

  n["cc"] = 3

  if maps.Equal(n, n2) {
    fmt.Println("n == n2")
  } else {
    fmt.Println("map not Equal")
  }

}


