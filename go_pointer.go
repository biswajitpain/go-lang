package main

import (
    "fmt"
)

func main(){
    x := new(int)
    y:=10
    x=&y
    fmt.Print(*x)
}
