// tag::hello[]
package main // <1>

import "fmt" // <2>
import (     // <3>
  "math"
  "time"
)

func hello() {
  fmt.Println("Hello World!")
  fmt.Println("sin(Pi) =", math.Sin(math.Pi))
}
// end::hello[]

// tag::const[]
const Pi = 3.14 // <1>
const ( // <2>
  // 将 1 左移 100 位来创建一个非常大的数字
  // 即这个数的二进制是 1 后面跟着 100 个 0
  Big = 1 << 100
  // 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
  Small = Big >> 99
)

func const_examples() {
  const World = "世界" // <3>
  fmt.Println("Hello", World, Pi)

  const Truth = true
  fmt.Println("Go rules?", Truth)
}
// end::const[]

func for_examples() {
// tag::for-1[]
  for i := 0; i < 5; i++ {
    fmt.Println("basic for loop: i =", i)
  }
// end::for-1[]

// tag::for-2[]
  s := 5
  for ; s < 5; {
    fmt.Println("for loop: s =", s)
  }
// end::for-2[]

// tag::for-3[]
  x := 0
  for x < 5 {
    fmt.Println("while x =", x)
    x += 1
  }
// end::for-3[]

// tag::for-4[]
  for { // <1>
    break // <2>
  }
// end::for-4[]
}

func if_examples() {
  x := 1.0

// tag::if-1[]
  if x > 0 {
    fmt.Println("if x > 0")
  }
// end::if-1[]

// tag::if-2[]
  if v := 1; v > 0 {
    fmt.Println("short if statement")
  }
// end::if-2[]

// tag::if-3[]
  if x > 0 {
    fmt.Println("x > 0")
  } else {
    fmt.Println("x <= 0")
  }
// end::if-3[]
}

func switch_examples() {
// tag::switch-1[]
  fmt.Printf("Today is ")
  switch time.Now().Weekday() {
  case time.Monday:
    fmt.Println("Mon")
  case time.Tuesday:
    fmt.Println("Tue")
  case time.Wednesday:
    fmt.Println("Wed")
  case time.Thursday:
    fmt.Println("Thu")
  case time.Friday:
    fmt.Println("Fri")
  case time.Saturday:
    fmt.Println("Sat")
    fallthrough // <1>
  case time.Sunday:
    fmt.Println("Sun")
  }
// end::switch-1[]

// tag::switch-2[]
  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("Good morning!")
  case t.Hour() < 17:
    fmt.Println("Good afternoon.")
  default:
    fmt.Println("Good evening.")
  }
// end::switch-2[]
}

func defer_examples() {
// tag::defer[]
  defer fmt.Println("defer")
  fmt.Printf("hello ")
// end::defer[]
}
