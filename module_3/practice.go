package main

import (
    "fmt"
)

func main () {
    go fmt.Printf("Hello?\n")

    fmt.Printf("World!\n")
    return
}
