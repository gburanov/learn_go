package main

import (
  "testing"
)

func TestPrimeNumbersUntil(t *testing.T) {
  if IsPrime(1) != true {
    t.Fatalf("1")
  }

  if IsPrime(2) != true {
    t.Fatalf("2")
  }

  if IsPrime(3) != true {
    t.Fatalf("3")
  }

  if IsPrime(4) == true {
    t.Fatalf("4")
  }
}
