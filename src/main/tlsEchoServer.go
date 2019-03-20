/* TLSEchoServer
 */
package main

import (
    "crypto/rand"
    "crypto/tls"
    "fmt"
    "time"
)

func tlsEchoServer() {

    cert, err := tls.LoadX509KeyPair("jan.newmarch.name.pem", "private.pem")
    checkError(err)
    config := tls.Config{Certificates: []tls.Certificate{cert}}

    now := time.Now()
    config.Time = func() time.Time { return now }
    config.Rand = rand.Reader

    service := "0.0.0.0:1200"

    listener, err := tls.Listen("tcp", service, &config)
    checkError(err)
    fmt.Println("Listening on ", service)

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println(err.Error())
            continue
        }
        fmt.Println("Accepted")
        go handleClient(conn)
    }
}
