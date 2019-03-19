/* DecryptingByRSA
 */

package main

import (
    "crypto/rsa"
    "crypto/rand"
    "fmt"
    "os"
    "encoding/gob"
)

func decryptingByRSA() {
    var key rsa.PrivateKey
    loadKey("private.key", &key)

    var message []byte
    loadMessage("encryByRSA", &message)

    // crypto/rand.Reader is a good source of entropy for randomizing the
    // encryption function.
    rng := rand.Reader

    secretMessage, err1 := rsa.DecryptPKCS1v15(rng, &key, message)
    checkError(err1)

    fmt.Println("The secret message is \"", string(secretMessage), "\"")

}

func loadMessage(fileName string, message interface{}) {
    inFile, err := os.Open(fileName)
    checkError(err)
    decoder := gob.NewDecoder(inFile)
    err = decoder.Decode(message)
    checkError(err)
    inFile.Close()
}
