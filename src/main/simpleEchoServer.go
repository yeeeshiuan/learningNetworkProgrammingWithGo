/* SimpleEchoServer
 */
package main

import (
    "net"
    "fmt"
)

func simpleEchoServer() {

    service := ":1201"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)

    fmt.Println("The port is ", service)

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        handleClientV1(conn)
        conn.Close() // we're finished
    }
}

func handleClientV1(conn net.Conn) {
    var buf [512]byte
    for {
        n, err := conn.Read(buf[0:])
        if err != nil {
            return
        }
        fmt.Println(string(buf[0:]))
        _, err2 := conn.Write(buf[0:n])
        if err2 != nil {
            return
        }
    }
}
