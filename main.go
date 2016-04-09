package main

import (
  "fmt"
)

func showStats(arr []int) {
  fmt.Println("Size", len(arr))
  fmt.Println("Capacity", cap(arr))
}

func main() {
  arr := []int{1, 2}
  showStats(arr)
  arr = append(arr, 3)
  showStats(arr)
}
