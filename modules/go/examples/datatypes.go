package main

import "fmt"
import "strings"

func pointer_examples() {
// tag::pointer[]
  i := 42

  // 类型 *T 是指向 T 类型值的指针。其零值为 nil。
  var p *int

  // & 操作符会生成一个指向其操作数的指针。
  p = &i

  // * 操作符表示指针指向的底层值。
  fmt.Println("*p =", *p) // 通过指针 p 读取 i

  // * 操作符表示指针
  *p = 21         // 通过指针 p 设置 i
  fmt.Println("*p =", *p, p)
// end::pointer[]
}

// tag::struct[]
type Vertex struct { // <1>
	X int // <2>
	Y int
}

func struct_examples() {
  var (
    v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
    v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
    v3 = Vertex{}      // X:0 Y:0
    p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
  )

  v1.X = 4
  p.Y = 5  // <3>
  fmt.Println("v1 =", v1)
  fmt.Println("v2 =", v2, "v2.X =", v2.X)
  fmt.Println("v3 =", v3, "v3.Y =", v3.Y)
  fmt.Println("p =", p, "*p =", *p)
}
// end::struct[]

func array_examples() {
// tag::array[]
  var a [2]int
  a[0] = 10
  a[1] = 20
  fmt.Println("array =", a)

  primes := [6]int{2, 3, 5, 7, 11, 13}
  fmt.Println("primes =", primes)
// end::array[]
}

// tag::print-slice[]
func printSlice(variable string, x []int) {
	fmt.Printf("%v len=%d cap=%d %v\n", variable, len(x), cap(x), x)
}
// end::print-slice[]

func slice_examples_1() {
// tag::slice-1[]
  primes := [6]int{2, 3, 5, 7, 11, 13}
  var s []int = primes[1:4] // <1>
  fmt.Println("slice of primes[1:4] =", s)

  q := []int{2, 3, 5, 7, 11, 13}
  fmt.Println("slice without array =", q)

  ss := []struct {
    i int
    b bool
  } {
    {1, true},
    {2, false},
    {3, true},
  }
  fmt.Println("slice of struct =", ss)
// end::slice-1[]

// tag::slice-2[]
  names := [4]string{
    "John",
    "Paul",
    "George",
    "Ringo",
  }
  fmt.Println("names =", names)

  name_of_slice := names[0:3]
  name_of_slice[0] = "XXX"
  fmt.Println("name_of_slice =", name_of_slice)
  fmt.Println("names =", names)
// end::slice-2[]

// tag::slice-3[]
  // 对于数组
  var a [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  // 以下切片是等价的：
  fmt.Println("a[0:10] =", a[0:10])
  fmt.Println("a[:10] =", a[:10])
  fmt.Println("a[0:] =", a[:10])
  fmt.Println("a[:] =", a[:])
// end::slice-3[]
}

func slice_examples_2() {
// tag::slice-4[]
  s := []int{2, 3, 5, 7, 11, 13}
  printSlice("s", s)

  // 截取切片使其长度为 0
  s = s[:0]
  printSlice("s", s)

  // 拓展其长度
  s = s[:4]
  printSlice("s", s)

  // 舍弃前两个值
  s = s[2:]
  printSlice("s", s)
// end::slice-4[]
}

func slice_examples_3() {
// tag::slice-5[]
  var s []int // <1>
  fmt.Println(s, len(s), cap(s))
  if s == nil {
    fmt.Println("s is nil!")
  }
// end::slice-5[]
}

func slice_examples_4() {
// tag::slice-6[]
a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
// end::slice-6[]
}

func slice_examples_5() {
// tag::slice-7[]
  // 创建一个井字板（经典游戏）
  board := [][]string{
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
  }

  // 两个玩家轮流打上 X 和 O
  board[0][0] = "X"
  board[2][2] = "O"
  board[1][2] = "X"
  board[1][0] = "O"
  board[0][2] = "X"

  for i := 0; i < len(board); i++ {
    fmt.Printf("%s\n", strings.Join(board[i], " "))
  }
// end::slice-7[]
}

func slice_examples_5() {
// tag::slice-8[]
  var s []int
  printSlice("s", s)

  // 添加一个空切片
  s = append(s, 0)
  printSlice("s", s)

  // 这个切片会按需增长
  s = append(s, 1)
  printSlice("s", s)

  // 可以一次性添加多个元素
  s = append(s, 2, 3, 4)
  printSlice("s", s)
// end::slice-8[]
}