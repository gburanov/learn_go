package main

import (
  "testing"
)

func TestPalindromes(t *testing.T) {
  if CheckIfPalindrome("abba") == false {
    t.Fatalf("abba")
  }

  if CheckIfPalindrome("abbat") == true {
    t.Fatalf("abbat")
  }
}
