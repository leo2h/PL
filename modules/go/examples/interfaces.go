package main

import "fmt"
import "math"

// tag::interface-1[]
type Abser interface {
	Abs() float64
}

type FloatA float64

func (f FloatA) Abs() float64 {
	if f < 0 {
		return -float64(f)
	}
	return float64(f)
}

type VertexA struct {
	X, Y float64
}

func (v *VertexA) Abs() float64 {
	if v == nil {
		return 0
	}
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func interface_examples_1() {
	var a Abser
	var v0 VertexA
	f := FloatA(-math.Sqrt2)
	v := VertexA{3.0, 4.0}

	a = f  // OK
	a = &v // OK

	// a = v is error
	fmt.Println("Call Abs of Abser", a.Abs())

	a = &v0 // <1>
	fmt.Println("Call Abs of Abser", a.Abs())
}
// end::interface-1[]

func interface_examples_2() {
// tag::interface-2[]
	var i interface{} // <1>
	fmt.Printf("(%v, %T)\n", i, i)
	i = 42      // <2>
	fmt.Printf("(%v, %T)\n", i, i)
	i = "hello" // <3>
	fmt.Printf("(%v, %T)\n", i, i)
// end::interface-2[]
}

func interface_examples_3() {
// tag::interface-3[]
    var i interface{} = "hello"

    s := i.(string)
    fmt.Println("string", s)

	s, ok := i.(string)
	fmt.Println("string", s, "ok", ok)

	f, ok := i.(float64)
	fmt.Println("float64", f, "ok", ok)

	// f = i.(float64) // panic
	// fmt.Println(f)
// end::interface-3[]
}

func interface_examples_4() {
// tag::interface-4[]
    var i interface{} = 42
    switch v := i.(type) {
    case int:
        fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
        fmt.Printf("%q is %v bytes long\n", v, len(v))
    default:
        fmt.Printf("I don't know about type %T!\n", v)
    }
// end::interface-4[]
}

// tag::interface-5[]
type Person struct {
    Name string
    Age  int
}

func (p *Person) String() string {
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func interface_examples_5() {
    a := Person{"Arthur Dent", 42}
    z := Person{"Zaphod Beeblebrox", 9001}
    fmt.Println("Person", a)
    fmt.Println("Person", z)
}
// end::interface-5[]
