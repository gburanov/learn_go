package main

import "fmt"

type Crier interface {
  Cry() string
}

type Cat struct {
}
 
func (_ Cat) Cry() string {
  return "Miow"
}

type Dog struct {
}

func (_ Dog) Cry() string {
  return "Bark"
}

func main() {
  var animal Crier
  animal = Cat{}
  fmt.Println(animal.Cry())
}
