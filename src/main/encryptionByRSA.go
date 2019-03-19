/* EncryptionByRSA
 * under
 * https://godoc.org/crypto/rsa#example-EncryptOAEP
 */

package main

import (
    "crypto/rand"
    "crypto/rsa"
    "fmt"
    "os"
    "encoding/gob"
)

func encryptionByRSA() {
    var publicKey rsa.PublicKey
    // from loadRSAKeys.go
    loadKey("public.key", &publicKey)

    fmt.Println("Public key modulus", publicKey.N.String())
    fmt.Println("Public key exponent", publicKey.E)

    secretMessage := []byte("PTT OGC")

    // crypto/rand.Reader is a good source of entropy for randomizing the
    // encryption function.
    rng := rand.Reader
    
    ciphertext, err := rsa.EncryptPKCS1v15(rng, &publicKey, secretMessage)
    checkError(err)

    // Since encryption is a randomized function, ciphertext will be
    // different each time.
    fmt.Printf("Ciphertext: %x\n", ciphertext)

    saveMessage("encryByRSA", &ciphertext)
}

func saveMessage(fileName string, message interface{}) {
    outFile, err := os.Create(fileName)
    checkError(err)
    encoder := gob.NewEncoder(outFile)
    err = encoder.Encode(message)
    checkError(err)
    outFile.Close()
}
