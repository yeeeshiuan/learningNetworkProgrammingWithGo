/* LookupPort
 */

package main

import (
    "net"
    "os"
    "fmt"
)

func lookupport(networkType, service string) {

    port, err := net.LookupPort(networkType, service)
    if err != nil {
        fmt.Println("Error: ", err.Error())
        os.Exit(2)
    }

    fmt.Println("Service port ", port)
}
