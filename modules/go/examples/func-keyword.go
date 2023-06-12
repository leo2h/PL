package main

import "fmt"

// tag::func[]
func add(x int, y int) int {
  return x + y
}

func sub(x, y int) int {
  return x - y
}

func swap(a, b string) (string, string) {
  return b, a // 多值返回
}
  
func div(x, y uint) (z uint, r uint) { // 命名返回值
  z = x / y
  r = x - z * y
  return
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