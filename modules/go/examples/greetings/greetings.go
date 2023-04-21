package greetings

import (
  "errors"
  "fmt"
  "math/rand"
  "time"
)

func init() {
  rand.Seed(time.Now().UnixNano())
}

func Hello(name string) (string, error) {
  // If no name was given, return an error with a message.
  if name == "" {
    return "", errors.New("empty name")
  }

  formats := []string {
    "Hi, %v. Welcome!",
    "Great to see you, %v!",
    "Hail, %v! Well met!",
  }

  message := fmt.Sprintf(formats[rand.Intn(len(formats))], name)
  return message, nil
}
