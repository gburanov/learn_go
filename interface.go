package main

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
