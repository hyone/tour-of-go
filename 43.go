// prepare libraries:
// $ export $GOPATH=/path/to/gopath
// $ go get code.google.com/p/go-tour/gotour

package main

import (
  "code.google.com/p/go-tour/wc"
  "strings"
)

func WordCount(s string) map[string]int {
  m := make(map[string]int)
  for _, v := range strings.Fields(s) {
    m[v]++
  }
  return m
}

func main () {
  wc.Test(WordCount)
}
