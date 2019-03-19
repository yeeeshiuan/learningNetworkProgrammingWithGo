/* UTF16 Server
 */
package main

import (
    "fmt"
    "net"
    "unicode/utf16"
)

func utf16Server() {

    service := "0.0.0.0:1210"
    tcpAddr, err := net.ResolveTCPAddr("tcp", service)
    checkError(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)

    fmt.Println("The port is ", service)

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }

        str := "j'ai arrêté"
        shorts := utf16.Encode([]rune(str))
        writeShorts(conn, shorts)

        conn.Close() // we're finished
    }
}

func writeShorts(conn net.Conn, shorts []uint16) {
    var bytes [2]byte

    // send the BOM as first two bytes
    bytes[0] = BOM >> 8
    bytes[1] = BOM & 255
    _, err := conn.Write(bytes[0:])
    if err != nil {
        return
    }

    for _, v := range shorts {
        bytes[0] = byte(v >> 8)
        bytes[1] = byte(v & 255)

        _, err = conn.Write(bytes[0:])
        if err != nil {
            return
        }
    }
}
