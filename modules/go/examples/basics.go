// tag::hello[]
package main // <1>

import "fmt" // <2>
import (     // <3>
  "time"
  "math"
  "math/rand"
)

func hello() {
  fmt.Println("Hello World!")
  fmt.Println("randint =", rand.Intn(10))
}
// end::hello[]

// tag::const[]
const Pi = 3.14 // <1>
const (         // <2>
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

// tag::func-1[]
func swap(x, y string) (string, string) {
  // 多值返回
	return y, x // <1>
}

func Mul(x, y int) (r int) { // <2>
  r = x * y
  return // <3>
}
// end::func-1[]

// tag::func-2[]
func compute(fn func(float64, float64) float64) float64 {
  return fn(3, 4)
}

func func_examples_2() {
  hypot := func(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
  }
  fmt.Println("hypot", hypot(5, 12))
  fmt.Println("compute", compute(hypot))
  fmt.Println("compute", compute(math.Pow))
}
// end::func-2[]

// tag::func-3[]
func adder() func(int) int {
  sum := 0
  return func(x int) int {
    sum += x
    return sum
  }
}

func func_examples_3() {
  pos, neg := adder(), adder()
  for i := 0; i < 10; i++ {
    fmt.Println(
      pos(i),
      neg(-2*i),
    )
  }
}
// end::func-3[]


// tag::var-1[]
var c, python, java bool
func var_examples_1() {
	var i int
	fmt.Println(i, c, python, java)
}
// end::var-1[]

func var_examples_2() {
// tag::var-2[]
  var i, j int = 1, 2
  fmt.Println(i, j)

  var c, python, java = true, false, "no!"
  fmt.Println(i, j, c, python, java)
// end::var-2[]
}

func var_examples_3() {
// tag::var-3[]
  k := 3
  c, python, java := true, false, "no!"
  fmt.Println(k, c, python, java)
// end::var-3[]

// end::var-4[]
  var (
    a int    // <1>
    b bool   // <2>
    s string // <3>
  )
  fmt.Println(a, b, s)
// end::var-4[]

// tag::var-5[]
  var v int = 1
  var f float64 = float64(v)
  var u uint = uint(f)
  fmt.Println(f, u)
// end::var-5[]

// tag::var-6[]
  v1 := 42 // 修改这里！
  fmt.Printf("v is of type %T\n", v1)
// end::var-6[]
}

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

func range_examples() {
// tag::range[]
  var pow_vals = []int{1, 2, 4, 8, 16, 32, 64, 128}
  for i, v := range pow_vals {
    fmt.Printf("2**%d = %d\n", i, v)
  }
  for _, v := range pow_vals {
    fmt.Printf("values = %d\n", v)
  }
  for i := range pow_vals { // <1>
    fmt.Printf("2**%d\n", i)
  }
// end::range[]
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

// tag::defer[]
func defer_examples() {
  defer fmt.Println("defer")
  fmt.Printf("hello ")
}
// end::defer[]
