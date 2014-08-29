// prepare libraries:
// $ export $GOPATH=/path/to/gopath
// $ go get code.google.com/p/go-tour/gotour

package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
  square := make([][]uint8, dy)
  for y := range square {
    square[y] = make([]uint8, dx)
    for x := range square[y] {
      square[y][x] = uint8(x*x + y*y)
    }
  }
  return square
}

func main () {
  pic.Show(Pic)
}
