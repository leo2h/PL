package main

import "fmt"

func main() {
  fmt.Println(">>>>>> hello <<<<<<")
  hello()

  fmt.Println(">>>>>> `func` keyword <<<<<<")
  h, w := swap("Hello", "World")
  fmt.Println(h, w, Mul(1, 1))

  fmt.Println("\n>>>>>> `var` keyword <<<<<<")
  var_examples_1()
  var_examples_2()

  fmt.Println("\n>>>>>> `const` keyword <<<<<<")
  const_examples()

  fmt.Println("\n>>>>>> `for` keyword <<<<<<")
  for_examples()

  fmt.Println("\n>>>>>> `range` keyword <<<<<<")
  range_examples()

  fmt.Println("\n>>>>>> `if` keyword <<<<<<")
  if_examples()

  fmt.Println("\n>>>>>> `defer` keyword <<<<<<")
  defer_examples()

  fmt.Println("\n>>>>>> `switch` keyword <<<<<<")
  switch_examples()

  fmt.Println("\n>>>>>> `struct` keyword <<<<<<")
  struct_examples()

  fmt.Println("\n>>>>>> `array` syntax <<<<<<")
  array_examples()

  fmt.Println("\n>>>>>> `slice` syntax <<<<<<")
  slice_examples_1()
  slice_examples_2()
  slice_examples_3()
  slice_examples_4()
  slice_examples_5()
  slice_examples_6()

  fmt.Println("\n>>>>>> `map` syntax <<<<<<")
  map_examples_1()

  fmt.Println("\n>>>>>> methods syntax <<<<<<")
  methods_examples_1()
  methods_examples_2()
  methods_examples_3()

  fmt.Println("\n>>>>>> interface syntax <<<<<<")
  interface_examples_1()
  interface_examples_2()
  interface_examples_3()
}
