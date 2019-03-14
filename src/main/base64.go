/**
 * Base64
 */

package main

import (
    "bytes"
    "encoding/base64"
    "fmt"
)

func base_64() {

    eightBitData := []byte{'1', '9', 'a', 'z', 'A', 'Z', 1, 128}
    fmt.Println("bytes: ", eightBitData)

    bb := &bytes.Buffer{}
    encoder := base64.NewEncoder(base64.StdEncoding, bb)
    encoder.Write(eightBitData)
    encoder.Close()
    fmt.Println("encode: ",bb)

    dbuf := make([]byte, 12) // why 12 bytes? sometimes, the last tow byte are 0.
    decoder := base64.NewDecoder(base64.StdEncoding, bb)
    decoder.Read(dbuf)

    fmt.Print("decode: ")
    for _, ch := range dbuf {
        fmt.Print(ch)
    }
    fmt.Println()
}
