/* UDPDaytimeClient
 */
package main

import (
    "net"
    "fmt"
)

func udpDaytimeClient(service string) {

    udpAddr, err := net.ResolveUDPAddr("udp4", service)
    checkError(err)

    conn, err := net.DialUDP("udp", nil, udpAddr)
    checkError(err)

    _, err = conn.Write([]byte("anything"))
    checkError(err)

    var buf [512]byte
    n, err := conn.Read(buf[0:])
    checkError(err)

    fmt.Println(string(buf[0:n]))
}
