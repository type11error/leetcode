package main

import (
  "fmt"
)

func isSubsequence(s string, t string) bool {
  i,j := 0,0
  for i < len(t) && j < len(s) {
    if t[i] == s[j] {
            j++
    }

    i++
  }

    return j == len(s)
}

func main() {
  res := isSubsequence("abc", "abjkc")
  fmt.Printf("result=%s", res)
}
