package main

import "fmt"

func main() {
    // Defer the call to the recover function
    msg := handlePanic()

    // Simulate a panic
    causePanic()

    // Continue with the normal flow of the program
    fmt.Println(msg)
}

func causePanic() {
    // The panic function is called with a message
    panic("Something went wrong!")
}

func handlePanic() (msg string) {
    // Defer a function to recover from the panic
    defer func() {
        if r := recover(); r != nil {
            // Print or log the panic details
            fmt.Println("Recovered from panic:", r)
            msg = "Program continues after panic recovery"
        }
    }()

    // Return the message to be printed
    return "HERE"
}
