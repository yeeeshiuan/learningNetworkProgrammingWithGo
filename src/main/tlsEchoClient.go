/* TLSEchoClient
 */
package main

import (
    "fmt"
    "crypto/tls"
)

func tlsEchoClient(service string) {
    conn, err := tls.Dial("tcp", service, nil)
    checkError(err)

    for n := 0; n < 10; n++ {
        fmt.Println("Writing...")
        conn.Write([]byte("Hello " + string(n+48)))

        var buf [512]byte
        n, err := conn.Read(buf[0:])
        checkError(err)

        fmt.Println(string(buf[0:n]))
    }
}
