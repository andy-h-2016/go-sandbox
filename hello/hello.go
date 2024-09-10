package main

import (
    "fmt"

    "github.com/andy-h-2016/go-sandbox/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
