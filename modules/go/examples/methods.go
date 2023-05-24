package main

import "fmt"
import "math"

// tag::method-1[]
type Vertexf struct {
  X, Y float64
}

func (v Vertexf) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func methods_examples_1() {
  v := Vertexf{3.0, 4.0}
  fmt.Printf("Abs(V%v) = %v\n", v, v.Abs())
}
// end::method-1[]

// tag::method-2[]
type MyFloat float64

func (f MyFloat) Abs() float64 {
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}

func methods_examples_2() {
  f := MyFloat(-math.Sqrt2)
  fmt.Println("Abs(MyFloat) =", f.Abs())
}
// end::method-2[]

// tag::method-3[]
func (v *Vertexf) Scale(factor float64) {
	v.X = v.X * factor
	v.Y = v.Y * factor
}

func methods_examples_3() {
  v := Vertexf{3.0, 4.0}
  p := &v
  p.Scale(5)
  v.Scale(10) // <1>
  fmt.Println("Abs Scaled Vertex =", v.Abs())
  fmt.Println("Abs Scaled Vertex =", p.Abs()) // <2>
}
// end::method-3[]
