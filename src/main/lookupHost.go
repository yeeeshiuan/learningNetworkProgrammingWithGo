/* LookupHost
 */

package main

import (
    "net"
    "os"
    "fmt"
)

func lookuphost(domain_name string) {

    name := domain_name

    addrs, err := net.LookupHost(name)
    if err != nil {
        fmt.Println("Error: ", err.Error())
        os.Exit(2)
    }

    for _, s := range addrs {
        fmt.Println(s)
    }
}

