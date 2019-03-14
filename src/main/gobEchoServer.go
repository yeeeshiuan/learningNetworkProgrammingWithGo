/* Gob EchoServer
 */
package main

import (
    "fmt"
    "net"
    "encoding/gob"
)

func gobEchoServer() {

    service := "0.0.0.0:1200"
    tcpAddr, err := net.ResolveTCPAddr("tcp", service)
    checkError(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)

    fmt.Println("The host:port is ", service)

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }

        encoder := gob.NewEncoder(conn)
        decoder := gob.NewDecoder(conn)

        for n := 0; n < 10; n++ {
            var person Person
            decoder.Decode(&person)
            fmt.Println(person.String())
            encoder.Encode(person)
        }
        conn.Close() // we're finished
    }
}
