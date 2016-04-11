package main

import (
  "strings"
  "unicode"
)

func Capitalize(str string) string {
  a := []rune(str)
  a[0] = unicode.ToUpper(a[0])
  str = string(a)
  return str
}

func UpCaseStrings(str string) string {
  words := strings.Fields(str)
  for _, word := range words {
    if len(word) < 4 {
      continue
    }
    str = strings.Replace(str, word, Capitalize(word), 1)
  }
  return str
}
