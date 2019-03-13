/* JSON EchoClient
 */
package main

import (
    "fmt"
    "net"
    "encoding/json"
)

func jsonEchoClient(service string) {
    person := Person{
        Name: Name{Family: "Newmarch", Personal: "Jan"},
        Email: []Email{Email{Kind: "home", Address: "jan@newmarch.name"},
            Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}

    conn, err := net.Dial("tcp", service)
    checkError(err)

    encoder := json.NewEncoder(conn)
    decoder := json.NewDecoder(conn)

    for n := 0; n < 10; n++ {
        encoder.Encode(person)
        var newPerson Person
        decoder.Decode(&newPerson)
        fmt.Println(newPerson.String())
    }
}
