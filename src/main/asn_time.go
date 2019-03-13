/* ASN.1
 */

package main

import (
    "encoding/asn1"
    "fmt"
    "time"
)

func asn_time() {
    t := time.Now()
    mdata, err := asn1.Marshal(t)
    checkError(err)

    var newtime time.Time
    _, err1 := asn1.Unmarshal(mdata, &newtime)
    checkError(err1)

    fmt.Println("After marshal/unmarshal: ", newtime)
}
