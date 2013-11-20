package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := 1.0 // Starting point.
    const n = 1000000 // Repeat n times.
    
    for i := 0; i < n; i++ {
        z = z - ((z * z - x) / (2 * z))
    }
    
    return z
}

func main() {
    var i float64
    fmt.Print("Input random number: ")
    fmt.Scan(&i)
    fmt.Println("Newton's method:", Sqrt(i))
    fmt.Println("math.Sqrt:", math.Sqrt(i))
}