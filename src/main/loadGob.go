/* LoadGob
 */

package main

import (
    "fmt"
    "os"
    "encoding/gob"
)

func loadgob() {
    var person Person
    loadGob("person.gob", &person)

    fmt.Println("Person", person.String())
}

func loadGob(fileName string, key interface{}) {
    inFile, err := os.Open(fileName)
    checkError(err)
    decoder := gob.NewDecoder(inFile)
    err = decoder.Decode(key)
    checkError(err)
    inFile.Close()
}
