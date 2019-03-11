/* main */

package main

import (
    "os"
    "fmt"
    "strings"
)

func main() {
    if len(os.Args) == 3 {
        switch strings.ToLower(os.Args[1]) {
            case "ip": 
                parse_ip(os.Args[2])
            case "mask":
                mask_ip(os.Args[2])
            case "resolve":
                resolve_ip(os.Args[2])
            case "lookuphost":
                lookuphost(os.Args[2])
            case "lookupc":
                lookupcanonical(os.Args[2])
            case "getheadinfo":
                getHeadInfo(os.Args[2])
            default:
                fmt.Printf("The choosing service(%s) is not supported.\n", os.Args[1]);
        }
    } else if len(os.Args) == 4 {
        switch strings.ToLower(os.Args[1]) {
            case "lookupport": 
                lookupport(os.Args[2], os.Args[3])
            default:
                fmt.Printf("The choosing service(%s) is not supported.\n", os.Args[1]);
        }
    } else {
        fmt.Fprintf(os.Stderr, "Usage: %s \"ip|mask\" ", os.Args[0])
        fmt.Fprintf(os.Stderr, "\"ip-address\"\n")

        fmt.Fprintf(os.Stderr, "Usage: %s \"resolve|lookupHost|lookupC\" ", os.Args[0])
        fmt.Fprintf(os.Stderr, "\"domain-name\"\n")

        fmt.Fprintf(os.Stderr, "Usage: %s \"getHeadInfo\" ", os.Args[0])
        fmt.Fprintf(os.Stderr, "\"host:port\"\n")

        fmt.Fprintf(os.Stderr, "Usage: %s \"lookupport\" \"network-type\" ", os.Args[0])
        fmt.Fprintf(os.Stderr, "\"service\"\n")
        os.Exit(1)
    }

    os.Exit(0)
}
