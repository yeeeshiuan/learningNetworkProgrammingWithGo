/* Blowfish
 */

package main

import (
    "bytes"
    "golang.org/x/crypto/blowfish"
    "fmt"
)

func blowfishSecurity() {
    key := []byte("my key")
    cipher, err := blowfish.NewCipher(key)
    if err != nil {
        fmt.Println(err.Error())
    }
    src := []byte("hello\n\n\n")
    var enc [512]byte

    cipher.Encrypt(enc[0:], src)

    var decrypt [16]byte
    cipher.Decrypt(decrypt[0:], enc[0:])
    result := bytes.NewBuffer(nil)
    result.Write(decrypt[0:16])
    fmt.Println(string(result.Bytes()))
}
