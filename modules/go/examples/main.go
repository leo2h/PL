package main

import "fmt"

func main() {
  fmt.Println(">>>>>> hello <<<<<<")
  hello()

  fmt.Println("\n>>>>>> `const` keyword <<<<<<")
  const_examples()

  fmt.Println("\n>>>>>> `for` keyword <<<<<<")
  for_examples()

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
}
