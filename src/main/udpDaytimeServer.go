/* UDPDaytimeServer
 */
package main

import (
    "fmt"
    "net"
    "time"
)

func udpDaytimeServer() {

    service := ":1200"
    udpAddr, err := net.ResolveUDPAddr("udp4", service)
    checkError(err)

    conn, err := net.ListenUDP("udp", udpAddr)
    checkError(err)

    fmt.Println("The port is ", service)

    for {
        handleClientUDP(conn)
    }
}

func handleClientUDP(conn *net.UDPConn) {

    var buf [512]byte

    _, addr, err := conn.ReadFromUDP(buf[0:])
    if err != nil {
        return
    }

    daytime := time.Now().String()

    conn.WriteToUDP([]byte(daytime), addr)
    fmt.Println("write: ", daytime)
}
