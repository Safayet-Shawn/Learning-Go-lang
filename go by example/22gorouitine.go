// go run 22gorouitine.go
package main

import (
    "fmt"
    "time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

    f("direct")

    go f("goroutine")

    go func(msg string) {
        
    time.Sleep(5*time.Second)
        fmt.Println(msg)
    }("going")

    time.Sleep(5*time.Second)
    fmt.Println("done")
}