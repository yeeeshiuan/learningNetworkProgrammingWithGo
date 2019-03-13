/* ASN.1
 */

package main

import (
    "encoding/asn1"
    "fmt"
)

func asn_str(number string) {
    mdata, err2 := asn1.Marshal(number)
    checkError(err2)

    var n string
    _, err3 := asn1.Unmarshal(mdata, &n)
    checkError(err3)

    fmt.Println("After marshal/unmarshal: ", n)
}
