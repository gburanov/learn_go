package main

import (
  "testing"
)

func TestPalindromes(t *testing.T) {
  if CheckIfPalindrome("a") == true {
    t.Fatalf("a")
  }

  if CheckIfPalindrome("abba") == false {
    t.Fatalf("abba")
  }

  if CheckIfPalindrome("abbat") == true {
    t.Fatalf("abbat")
  }
}

func TestPalindromesCount(t *testing.T) {
  palindromes := CountPalindromes("barbarabar")
  t.Log(palindromes)
  if (len(palindromes) != 3) {
    t.Fatalf("Shoud be 3 but %d", len(palindromes))
  }
}
