package main

import (
  "fmt"
  // "math"
)

func Sqrt(x float64) float64 {
  z    := 1.0
  prev := 0.0
  // for math.Abs(z - prev) > 0 {
  for z != prev {
    prev = z
    z = z - (z * z - x) / (2 * x)
  }
  return z
}

func main () {
  fmt.Println(Sqrt(2))
}
