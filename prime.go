package main

// also try resheto
func IsPrime(num int) bool {
  if (num == 1 || num == 2) {
    return true
  }

  for i := 2; i < num; i++ {
    if num % i == 0 {
      return false
    }
  }
  return true
}

func PrimeNumbersUntil(final int) []int {
  ret := []int{}
  return ret
}
