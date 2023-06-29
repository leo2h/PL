package main

import "fmt"

// tag::func[]
func EmptyFunction() {
  
}

func AddI(a int, b int) int {
  return a + b
}

func AddF(a, b float32) float32 {
  return a + b
}
// end::func[]

func main()  {
  fmt.Println("1 + 1 =", add(1, 1))
  fmt.Println("2 - 1 =", sub(2, 1))

  z, r := div(10, 3)
  fmt.Println("10 / 3 =", z, "...", r)

  s1, s2 := swap("s1", "s2")
  fmt.Printf("swap(%v, %v) = %v, %v\n", "s1", "s2", s1, s2)
}