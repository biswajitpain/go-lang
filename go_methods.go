//Sample Program  To show how declare meethods and how the go complier supports them

package main

import (
        "fmt"
)

type user struct{
    name string
    email string
}


//Notify implements a methods with a value receiver.

func (u user) notify(){
    fmt.Printf("User: Sending user email to %s<%s>\n",u.name,u.email)
}

//change Email implements a methods with a pointer receiver.

func (u *user) changeEmail(email string){
    u.email = email
}


func main(){
    //values of type user can be used to call methods
    //Declare with a value receiver

    bill := user{"Biswajit" ,"bpain2010@gmail.com"}
    bill.notify()
    bill.changeEmail("biswajit.pain@olacabs.com")

    //Pointers of type user can also be used to methods
    //declared with a value receiver.

    poulami := &user{"Poulami" ,"saha.poulami91@gmail.com"}
    poulami.notify()
    poulami.changeEmail("poulami.saha@bpcse.in")

}

