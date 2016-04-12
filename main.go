package main

import (
  "fmt"
)

func printTheShit(val interface{}) {
  fmt.Println(val)
}

func showStats(arr []int) {
  fmt.Println("Size", len(arr))
  fmt.Println("Capacity", cap(arr))
}
