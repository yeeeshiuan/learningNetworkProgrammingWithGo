/* DaytimeServer
 */
package main

import (
    "net"
    "fmt"
    "time"
)

func daytimeServer() {

    service := ":1200"
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

        daytime := time.Now().String()
        conn.Write([]byte(daytime))     // don't care about return value
        fmt.Println("write: ", daytime)
        conn.Close()                    // we're finished with this client
    }
}
