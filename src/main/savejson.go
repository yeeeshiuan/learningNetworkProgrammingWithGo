/* SaveJSON
 */

package main

import (
    "encoding/json"
    "os"
)

func savejson() {
    person := Person{
        Name: Name{Family: "Newmarch", Personal: "Jan"},
        Email: []Email{Email{Kind: "home", Address: "jan@newmarch.name"},
            Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}

    saveJSON("person.json", person)
}

func saveJSON(fileName string, key interface{}) {
    outFile, err := os.Create(fileName)
    checkError(err)
    encoder := json.NewEncoder(outFile)
    err = encoder.Encode(key)
    checkError(err)
    outFile.Close()
}
