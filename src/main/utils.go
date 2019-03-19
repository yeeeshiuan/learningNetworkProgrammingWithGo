/* utils */

package main

import (
    "os"
    "fmt"
)

const BOM = '\ufffe'

type Person struct {
    Name  Name
    Email []Email
}

type Name struct {
    Family   string
    Personal string
}

type Email struct {
    Kind    string
    Address string
}

func (p Person) String() string {
    s := p.Name.Personal + " " + p.Name.Family
    for _, v := range p.Email {
        s += "\n" + v.Kind + ": " + v.Address
    }
    return s
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
        os.Exit(1)
    }
}
