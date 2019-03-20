/* utils */

package main

import (
    "os"
    "fmt"
    "encoding/gob"
    "encoding/pem"
    "crypto/x509"
    "crypto/rsa"
    "net"
)

const BOM = '\ufffe'

type Person struct {
    Name  Name
    Email []Email
}

type Name struct {
    Family   string
    Personal string
}

type Email struct {
    Kind    string
    Address string
}

func (p Person) String() string {
    s := p.Name.Personal + " " + p.Name.Family
    for _, v := range p.Email {
        s += "\n" + v.Kind + ": " + v.Address
    }
    return s
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
        os.Exit(1)
    }
}

func saveGobKey(fileName string, key interface{}) {
    outFile, err := os.Create(fileName)
    checkError(err)
    encoder := gob.NewEncoder(outFile)
    err = encoder.Encode(key)
    checkError(err)
    outFile.Close()
}

func savePEMKey(fileName string, key *rsa.PrivateKey) {

    outFile, err := os.Create(fileName)
    checkError(err)

    var privateKey = &pem.Block{Type: "RSA PRIVATE KEY",
        Bytes: x509.MarshalPKCS1PrivateKey(key)}

    pem.Encode(outFile, privateKey)

    outFile.Close()
}

func loadKey(fileName string, key interface{}) {
    inFile, err := os.Open(fileName)
    checkError(err)
    decoder := gob.NewDecoder(inFile)
    err = decoder.Decode(key)
    checkError(err)
    inFile.Close()
}

func handleClient(conn net.Conn) {
    defer conn.Close()

    var buf [512]byte
    for {
        fmt.Println("Trying to read")
        n, err := conn.Read(buf[0:])
        if err != nil {
            fmt.Println(err)
        }
        _, err2 := conn.Write(buf[0:n])
        if err2 != nil {
            return
        }
    }
}
