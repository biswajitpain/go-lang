package main

import "fmt"

func main(){
    x := 0
    mult := func(x, y int) int {
        return x*y
    }
    fmt.Println(mult(3,4))

    incr := func() int {
        x++
        return x
    }
    fmt.Println(incr())
    fmt.Println(incr())
}
