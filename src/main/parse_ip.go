/* IP */

package main

import (
    "net"
    "fmt"
)

func parse_ip(ip string) {
    name := ip

    addr := net.ParseIP(name)
    if addr == nil {
        fmt.Println("Invalid address")
    } else {
        fmt.Println("The address is ", addr.String())
    }
}
