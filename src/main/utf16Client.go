/* UTF16 Client
 */
package main

import (
    "fmt"
    "net"
    "unicode/utf16"
)

func utf16Client(service string) {

    conn, err := net.Dial("tcp", service)
    checkError(err)

    shorts := readShorts(conn)
    ints := utf16.Decode(shorts)
    str := string(ints)

    fmt.Println(str)
}

func readShorts(conn net.Conn) []uint16 {
    var buf [512]byte

    // read everything into the buffer
    n, err := conn.Read(buf[0:2])
    for true {
        m, err := conn.Read(buf[n:])
        if m == 0 || err != nil {
            break
        }
        n += m
    }

    checkError(err)
    var shorts []uint16
    shorts = make([]uint16, n/2)

    if buf[0] == 0xff && buf[1] == 0xfe {
        // big endian
        for i := 2; i < n; i += 2 {
            shorts[i/2] = uint16(buf[i])<<8 + uint16(buf[i+1])
        }
    } else if buf[1] == 0xfe && buf[0] == 0xff {
        // little endian
        for i := 2; i < n; i += 2 {
            shorts[i/2] = uint16(buf[i+1])<<8 + uint16(buf[i])
        }
    } else {
        // unknown byte order
        fmt.Println("Unknown order")
    }
    return shorts
}
