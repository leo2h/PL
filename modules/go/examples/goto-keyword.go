package main

import "fmt"

func main() {
  var s int = 0
  for i := 0; i < 10; i++ {
    s += i
    if s > 5 {
      goto end;
    }
  }

end:
  fmt.Println("end by goto")
}
