package main

func CheckIfPalindrome(str string) bool {
  len := len(str)
  if (len < 2) {
    return false
  }

  for i := 0; i < len/2; i++ {
    if str[i] != str[len - i - 1] {
      return false
    }
  }
  return true
}

func CountPalindromes(str string) []string {
  len := len(str)
  palindromes := []string{}
  for start := 0; start < len; start++ {
    for finish := len-1; finish > start; finish-- {
      check_str := str[start:finish]
      if (CheckIfPalindrome(check_str)) {
        palindromes = append(palindromes, check_str)
      }
    }
  }
  return palindromes
}
