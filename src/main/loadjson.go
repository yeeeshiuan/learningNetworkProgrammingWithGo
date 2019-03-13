/* LoadJSON
 */

package main

import (
    "encoding/json"
    "fmt"
    "os"
)

func loadjson() {
    var person Person
    loadJSON("person.json", &person)

    fmt.Println("Person", person.String())
}

func loadJSON(fileName string, key interface{}) {
    inFile, err := os.Open(fileName)
    checkError(err)
    decoder := json.NewDecoder(inFile)
    err = decoder.Decode(key)
    checkError(err)
    inFile.Close()
}
