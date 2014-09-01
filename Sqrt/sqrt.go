package main

import (
    "fmt"
)

func Sqrt(x float64) float64 {
    z := x
    next := float64(1)
    for z-next > 0.0000000000000005 {
    	z = z-(z*z-x)/(2*z)
    	next = z-(z*z-x)/(2*z)
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
}
