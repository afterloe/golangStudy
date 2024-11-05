package main

import (
    "fmt"
    "time"
    "strconv"
)

type MyError struct {
    When time.Time
    What string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
    return &MyError {
        time.Now(),
        "it didn't work",
    }
}

func main(){
    if err := run(); err != nil {
        // panic(err)
        fmt.Println(err)
    }

    i, err := strconv.Atoi("122")
    if nil != err {
        fmt.Printf("Couldn't convert numnber: %v \r\n", err)
    }
    fmt.Printf("Converted integer: %d \r\n", i)
}

