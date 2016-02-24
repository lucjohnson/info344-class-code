package main

import (
    "fmt"
    "time"
)

func main() {
    // when you know the length for sure, use an array
    // if the length may change (dynamic) use a slice
    // can generate a slice of an array by doing arrayName[x:y]
    var months [12]string
    
    for idx := 0; idx < 12; idx++ {
        months[idx] = time.Month(idx+1).String()
    }
    
    fmt.Println(months)
    
    // when creating a slice you don't give it a size
    var dynoMonths []string
    
    for idx := 0; idx < 12; idx++ {
        dynoMonths = append(dynoMonths, time.Month(idx+1).String())
    }
    
    fmt.Println(dynoMonths)
}