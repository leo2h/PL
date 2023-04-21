package main

import (
  "fmt"
  "log"

  "example.com/greetings"
)

func main() {
  log.SetPrefix("greetings: ")
  log.SetFlags(0)

  message, err := greetings.Hello("Leo")
  if err == nil {
    fmt.Println(message)
  }

  message, err = greetings.Hello("")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(message)
}
