package main

import (
        "os"
        "log"
        )
func check_err(e error){
    if e != nil {
        panic(e)
    }
}

const LOGFILE = "/tmp/logfile.txt"
func main(){
    fa, err := os.OpenFile(LOGFILE, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    check_err(err)
    defer fa.Close()
    log.SetOutput(fa)


    data := os.Args[1]
    log.Println(data)
    f, err := os.OpenFile("/tmp/data",os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    check_err(err)
    defer f.Close()
    ret , err := f.WriteString(data)
    f.WriteString("\n")

    check_err(err)
    log.Println("Wrote %d bytes", ret)

}
