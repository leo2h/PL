package main

import "fmt"
import "time"

func main() {
forever_loop:
  fmt.Println(time.Now())
  time.Sleep(2 * time.Second)
  goto forever_loop;
}
