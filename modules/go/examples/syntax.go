// tag::hello[]
// 每个 Go 程序都是由包构成的
// 程序从 main 包开始运行
package main

import "fmt" // <1>
import (     // <2>
  "math"
  "time"
)

func hello() {
  fmt.Println("Hello world!")
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

func const_statement() {
  const World = "世界" // <3>
  fmt.Println("Hello", World, Pi)

  const Truth = true
  fmt.Println("Go rules?", Truth)
}
// end::const[]

// tag::for[]
func for_statement() {
  // 基本的 for 循环由三部分组成，它们用分号隔开：
  // 1. 初始化语句：在第一次迭代前执行
  // 2. 条件表达式：在每次迭代前求值
  // 3. 后置语句：在每次迭代的结尾执行
  s1 := 0
  for i := 0; i < 10; i++ {
    s1 += 1
  }
  fmt.Println(s1)

  s2 := 0
  // 初始化语句和后置语句是可选的
  for ; s2 < 1000; {
    s2 += 1
  }
  fmt.Println(s2)

  s3 := 0
  // for 是 Go 中的 “while”
  for s3 < 1000 {
    s3 += 1
  }
  fmt.Println(s3)

  // 无限循环
  // 如果省略循环条件，该循环就不会结束
  // 因此无限循环可以写得很紧凑
  for {
    if s1 > 0 {
      break
    }
  }
}
// end::for[]

// tag::if[]
func if_statement(x float64) {
  if x < 0 {
    fmt.Println("x < 0")
  }

  // if 的简短语句
  // 同 for 一样， if 语句可以在条件表达式前执行一个简单的语句。
  // 该语句声明的变量作用域仅在 if 之内。
  if v := 0; v > 0 {
    fmt.Println("v > 0")
  }

  // if 和 else
  if x > 0 {
    fmt.Println("x > 0")
  } else {
    fmt.Println("x <= 0")
  }
}
// end::if[]

// tag::switch[]
func switch_statement() {
  wd := time.Now().Weekday()
  switch wd {
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
    fallthrough //
  case time.Sunday:
    fmt.Println("Sun")
  }

  // 没有条件的 switch
  // 没有条件的 switch 同 switch true 一样。
  // 这种形式能将一长串 if-then-else 写得更加清晰。
  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("Good morning!")
  case t.Hour() < 17:
    fmt.Println("Good afternoon.")
  default:
    fmt.Println("Good evening.")
  }
}
// end::switch[]

// tag::defer[]
func defer_statement() {
  defer fmt.Println("World")
  fmt.Println("Hello")
}
// end::defer[]

// tag::pointer[]
func pointer_statement() {
  i := 42

  // 类型 *T 是指向 T 类型值的指针。其零值为 nil。
  var p *int

  // & 操作符会生成一个指向其操作数的指针。
  p = &i

  // * 操作符表示指针指向的底层值。
  fmt.Println(*p) // 通过指针 p 读取 i
  *p = 21         // 通过指针 p 设置 i
}
// end::pointer[]

// tag::struct[]
type Vertex struct {
	X int
	Y int
}

func struct_statement() {
  var (
    v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
    v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
    v3 = Vertex{}      // X:0 Y:0
    p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
  )

  v := Vertex{1, 2}
  q := &v
  v.X = 4
  q.Y = 5  // 指针访问字段也是用 .
  fmt.Println(v, v.Y, v1, v2, v3, p, q)
}
// end::struct[]

// end::array[]
func array_statement() {
  var a [2]int
  a[0] = 10
  a[1] = 20

  primes := [6]int{2, 3, 5, 7, 11, 13}
  fmt.Println(a, primes)
}
// end::array[]

func main() {
  hello()
  const_statement()
  if_statement(10.0)
  for_statement()
  defer_statement()
  switch_statement()
  struct_statement()
  array_statement()

  fmt.Println(math.Pow(1, 1))
}
