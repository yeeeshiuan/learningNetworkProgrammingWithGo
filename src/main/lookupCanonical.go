/* LookupCanonical
 */

package main

import (
    "net"
    "os"
    "fmt"
)

func lookupcanonical(domain_name string) {

    name := domain_name

    cname, err := net.LookupCNAME(name)
    if err != nil {
        fmt.Println("Error: ", err.Error())
        os.Exit(2)
    }

    fmt.Println(cname)
}

