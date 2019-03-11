/* resolve_ip */

package main

import (
    "net"
    "os"
    "fmt"
    "reflect"
)

func resolve_ip(domain_name string) {

    name := domain_name
    // where net is one of "ip", "ip4" or "ip6"
    addr, err := net.ResolveIPAddr("ip4", name)
    if err != nil {
        fmt.Println("Resolution error", err.Error())
        os.Exit(1)
    }
    fmt.Println("Resolved address is ", addr.String())
    fmt.Println("type of addr is ", reflect.TypeOf(addr))
    fmt.Println("ip ", addr.IP)
    fmt.Println("type of addr.IP is ", reflect.TypeOf(addr.IP))
}
