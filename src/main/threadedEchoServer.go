/* threadedEchoServer
 */
package main

import (
    "net"
    "fmt"
)

func threadedEchoServer() {

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
        // run as a goroutine
        go handleClient(conn)
    }
}

func handleClient(conn net.Conn) {
    // close connection on exit
    defer conn.Close()

    var buf [512]byte
    for {
        // read upto 512 bytes
        n, err := conn.Read(buf[0:])
        if err != nil {
            return
        }
        fmt.Println(string(buf[0:]))
        // write the n bytes read
        _, err2 := conn.Write(buf[0:n])
        if err2 != nil {
            return
        }
    }
}
