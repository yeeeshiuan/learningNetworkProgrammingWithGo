/* ASN.1 DaytimeClient
 */
package main

import (
    "encoding/asn1"
    "fmt"
    "io/ioutil"
    "net"
    "time"
)

func asn_daytime_client(service string) {

    conn, err := net.Dial("tcp", service)
    checkError(err)

    defer conn.Close()

    result, err := ioutil.ReadAll(conn)
    checkError(err)

    var newtime time.Time
    _, err1 := asn1.Unmarshal(result, &newtime)
    checkError(err1)

    fmt.Println("After marshal/unmarshal: ", newtime.String())
}
