package main

import (
    "fmt"
    "sync"
)

func main () {
    var foo int = 0
    var wg *sync.WaitGroup = new(sync.WaitGroup)

    for i := 0; i < 5; i++ {
        wg.Add(2)
        go addToNumber(&foo, wg)
        go addToNumber(&foo, wg)
    }

    wg.Wait()
    return
}

func addToNumber (num *int, wg *sync.WaitGroup) {
    (*num)++

    fmt.Printf("Number value:\t%d\n", *num)

    wg.Done()
    return
}
