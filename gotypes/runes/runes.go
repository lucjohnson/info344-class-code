package main

import (
    "fmt"
)

func main() {
    s:= "Hello, 世界"
    
    // this is how to iterate over a string in go
    for idx, r := range s {
        fmt.Println(idx, r, string(r))
    }
}