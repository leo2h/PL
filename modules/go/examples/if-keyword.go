package main

import "fmt"
import "math"

// tag::if-1[]
func factorial(n uint) uint {
  if n > 0 {
    return n * factorial(n - 1)
  }
  return 1
}

func hello(name string) {
  if name == "" {
    fmt.Println("Hello someone.")
  } else {
    fmt.Println("Hello", name)
  }
}
// end::if-1[]

// tag::if-2[]
func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  }
  return lim
}
// end::if-2[]

func main() {
  hello("")
  hello("Bob")
  fmt.Println("pow(3, 2, 10) =", pow(3, 2, 10))
  fmt.Println("5! =", factorial(5))
}