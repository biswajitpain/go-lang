package main

import(
    "fmt"
    "flag"
)

func main(){


    ststusPtr := flag.String("status","OK","Default Status is ok")

    flag.Parse()

    fmt.Println("Status is", *ststusPtr)
    fmt.Println("tail:", flag.Args())
}


