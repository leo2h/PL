package main

import "fmt"

func main() {
  var i int
  var i8 int8
  var i16 int16
  var i32 int32
  var i64 int64

  var u uint
  var u8 uint8
  var u16 uint16
  var u32 uint32
  var u64 uint64
  var ptr uintptr

  var (
    f32 float32
    f64 float64
  )

  var c64 complex64
  var c128 complex128

  var b byte   // <1>
  var r rune   // <2>
  var s string

  fmt.Printf("zero of %8T = %v\n", i, i);
  fmt.Printf("zero of %8T = %v\n", i8, i8);
  fmt.Printf("zero of %8T = %v\n", i16, i16);
  fmt.Printf("zero of %8T = %v\n", i32, i32);
  fmt.Printf("zero of %8T = %v\n", i64, i64);

  fmt.Printf("zero of %8T = %v\n", u, u);
  fmt.Printf("zero of %8T = %v\n", u8, u8);
  fmt.Printf("zero of %8T = %v\n", u16, u16);
  fmt.Printf("zero of %8T = %v\n", u32, u32);
  fmt.Printf("zero of %8T = %v\n", u64, u64);

  fmt.Printf("zero of %8T = %v\n", f32, f32);
  fmt.Printf("zero of %8T = %v\n", f64, f64);

  fmt.Printf("zero of %8T = %v\n", ptr, ptr);
  fmt.Printf("zero of %8s = %v\n", "byte", b);
  fmt.Printf("zero of %8s = %v\n", "rune", r);
  fmt.Printf("zero of %8T = %v\n", s, s);

  fmt.Printf("zero of %10T = %v\n", c64, c64);
  fmt.Printf("zero of %10T = %v\n", c128, c128);
}