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

func main() {
  printTheShit("Start execution")

  arr := []int{1, 2}
  showStats(arr)
  arr = append(arr, 3)
  showStats(arr)
}
