package main

import "fmt"
import "sync"
import "time"

// tag::goroutine-1[]
func say(s string) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(s)
}

func goroutine_examples_1() {
	go say("world") // <1>
	say("hello")    // <2>
}
// end::goroutine-1[]

// tag::goroutine-2[]
func sum_routine(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // <1>
}

func goroutine_examples_2() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) // <2>
	go sum_routine(s[:len(s)/2], c)
	go sum_routine(s[len(s)/2:], c)
	x, y := <-c, <-c // <3>

	fmt.Println(x, y, x+y)
}
// end::goroutine-2[]

func goroutine_examples_3() {
// tag::goroutine-3[]
	ch := make(chan int, 2) // <1>
	ch <- 1
	ch <- 2
	fmt.Println("read from ch", <-ch, <-ch)
// end::goroutine-3[]
}

// tag::goroutine-4[]
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func goroutine_examples_4() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	fmt.Printf("fibonacci:")
	for v := range c {
		fmt.Printf(" %v", v)
	}
	fmt.Println(".")

	v, ok := <- c // <1>
	if ok {
		fmt.Println(v)
	}
}
// end::goroutine-4[]

// tag::goroutine-5[]
func fibonacci_1(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println(" quit")
			return
		}
	}
}

func goroutine_examples_5() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		fmt.Printf("fibonacci:")
		for i := 0; i < 10; i++ {
			fmt.Printf(" %v", <-c)
		}
		quit <- 0
	}()
	fibonacci_1(c, quit)
}
// end::goroutine-5[]

func goroutine_examples_6() {
// tag::goroutine-6[]
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.   ")
		case <-boom:
			fmt.Println("BOOM!   ")
			return
		default: // <1>
			fmt.Println("default.")
			time.Sleep(50 * time.Millisecond)
		}
	}
// end::goroutine-6[]
}

// tag::goroutine-7[]
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	defer c.mux.Unlock()
	return c.v[key]
}

func goroutine_examples_7() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println("count of somekey:", c.Value("somekey"))
}
// end::goroutine-7[]
