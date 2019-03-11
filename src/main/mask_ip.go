/* Mask
 */

package main

import (
    "fmt"
    "net"
)

func mask_ip(ip string) {
    dotAddr := ip

    addr := net.ParseIP(dotAddr)
    if addr == nil {
        fmt.Println("Invalid address")
    } else {
        mask := addr.DefaultMask()
        network := addr.Mask(mask)
        ones, bits := mask.Size()
        fmt.Println("Address is ", addr.String())
        fmt.Println("Default mask length is ", bits)
        fmt.Println("Leading ones count is ", ones)
        fmt.Println("Mask is (hex) ", mask.String())
        fmt.Println("Network is ", network.String())
    }
}
