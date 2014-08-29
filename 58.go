package main

import (
  "fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
  // a call to fmt.Print(e) inside the Error method will send the program into an infinite loop.
  // You can avoid this by converting e first: fmt.Print(float64(e)).
  return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// newton's method
func Sqrt(x float64) (float64, error) {
  if x < 0 {
    return x, ErrNegativeSqrt(x)
  }
  z    := 1.0
  prev := 0.0
  for z != prev {
    prev = z
    z = z - (z * z - x) / (2 * x)
  }
  return z, nil
}


func main () {
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(-2))
}
