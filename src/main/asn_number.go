/* ASN.1
 */

package main

import (
    "encoding/asn1"
    "fmt"
    "strconv"
)

func asn_number(number string) {
    i, err1 := strconv.Atoi(number)
    checkError(err1)

    mdata, err2 := asn1.Marshal(i)
    checkError(err2)

    var n int
    _, err3 := asn1.Unmarshal(mdata, &n)
    checkError(err3)

    fmt.Println("After marshal/unmarshal: ", n)
}
