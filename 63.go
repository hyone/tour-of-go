package main

import (
  "io"
  "os"
  "strings"
)

type rot13Reader struct {
  r io.Reader
}

func (reader *rot13Reader) Read(p []byte) (n int, err error) {
  n, err = reader.r.Read(p)
  if err != nil {
    return 0, err
  }

  for i := 0; i < n; i++ {
    p[i] = decodeROT13(p[i])
  }

  return n, nil
}


func decodeROT13(c byte) byte {
  var r byte

  switch {
  case 'A' <= c && c <= 'Z':
    r = (c - 'A' + 13) % 26 + 'A'
  case 'a' <= c && c <= 'z':
    r = (c - 'a' + 13) % 26 + 'a'
  default:
    r = c
  }

  return r
}


func main () {
  s := strings.NewReader("Lbh penpxrq gur pbqr!")
  r := rot13Reader { s }
  io.Copy(os.Stdout, &r)
}

