/* main */

package main

import (
    "os"
    "fmt"
    "strings"
)

func main() {
    if len(os.Args) == 2 {
        switch strings.ToLower(os.Args[1]) {
            case "help": 
                help()
            case "daytimeserver": 
                daytimeServer()
            case "simpleechoserver":
                simpleEchoServer()
            case "threadedechoserver":
                threadedEchoServer()
            case "udpdaytimeserver":
                udpDaytimeServer()
            case "asn_time":
                asn_time()
            case "asn_daytime_server":
                asn_daytime_server()
            case "savejson":
                savejson()
            case "loadjson":
                loadjson()
            case "jsonechoserver":
                jsonEchoServer()
            case "savegob":
                savegob()
            case "loadgob":
                loadgob()
            case "gobechoserver":
                gobEchoServer()
            case "base64":
                base_64()
            case "ftpserver":
                ftpServer()
            case "utf16server":
                utf16Server()
            case "md5hash":
                md5Hash()
            case "blowfishsecurity":
                blowfishSecurity()
            case "genrsakeys":
                genRSAKeys()
            case "loadrsakeys":
                loadRSAKeys()
            case "encryptionbyrsa":
                encryptionByRSA()
            case "decryptingbyrsa":
                decryptingByRSA()
            default:
                fmt.Printf("You wrong(%s).\n", os.Args[1])
        }
    }else if len(os.Args) == 3 {
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
            case "udpdaytimeclient":
                udpDaytimeClient(os.Args[2])
            case "ping":
                ping(os.Args[2])
            case "asn_number":
                asn_number(os.Args[2])
            case "asn_str":
                asn_str(os.Args[2])
            case "asn_daytime_client":
                asn_daytime_client(os.Args[2])
            case "jsonechoclient":
                jsonEchoClient(os.Args[2])
            case "gobechoclient":
                gobEchoClient(os.Args[2])
            case "ftpclient":
                ftpClient(os.Args[2])
            case "utf16client":
                utf16Client(os.Args[2])
            default:
                fmt.Printf("You wrong(%s).\n", os.Args[1])
        }
    } else if len(os.Args) == 4 {
        switch strings.ToLower(os.Args[1]) {
            case "lookupport": 
                lookupport(os.Args[2], os.Args[3])
            default:
                fmt.Printf("You wrong(%s).\n", os.Args[1])
        }
    } else {
        help()
        os.Exit(1)
    }

    os.Exit(0)
}

func help() {
    fmt.Fprintf(os.Stderr, "Usage: %s \"help|daytimeServer|simpleEchoServer\" \n", os.Args[0])
    fmt.Fprintf(os.Stderr, "Usage: %s \"threadedEchoServer|udpDaytimeServer\" \n", os.Args[0])
    fmt.Fprintf(os.Stderr, "Usage: %s \"asn_time|asn_daytime_server|saveJSON|loadJSON\" \n", os.Args[0])
    fmt.Fprintf(os.Stderr, "Usage: %s \"jsonEchoServer|saveGob|loadGob|gobEchoServer\" \n", os.Args[0])
    fmt.Fprintf(os.Stderr, "Usage: %s \"base64|ftpServer|utf16Server|md5Hash\" \n", os.Args[0])
    fmt.Fprintf(os.Stderr, "Usage: %s \"blowfishSecurity|genRSAKeys|loadRSAKeys\" \n", os.Args[0])
    fmt.Fprintf(os.Stderr, "Usage: %s \"encryptionByRSA|decryptingByRSA\" \n", os.Args[0])

    fmt.Fprintf(os.Stderr, "Usage: %s \"ip|mask\" ", os.Args[0])
    fmt.Fprintf(os.Stderr, "\"ip-address\"\n")

    fmt.Fprintf(os.Stderr, "Usage: %s \"ASN_number\" ", os.Args[0])
    fmt.Fprintf(os.Stderr, "\"number\"\n")

    fmt.Fprintf(os.Stderr, "Usage: %s \"ASN_str\" ", os.Args[0])
    fmt.Fprintf(os.Stderr, "\"string\"\n")

    fmt.Fprintf(os.Stderr, "Usage: %s \"resolve|lookupHost|lookupC|ftpClient\" ", os.Args[0])
    fmt.Fprintf(os.Stderr, "\"domain-name\"\n")

    fmt.Fprintf(os.Stderr, "Usage: sudo %s \"ping\" ", os.Args[0])
    fmt.Fprintf(os.Stderr, "\"localhost\"\n")

    fmt.Fprintf(os.Stderr, "Usage: %s \"getHeadInfo|udpDaytimeClient|asn_daytime_client\" ", os.Args[0])
    fmt.Fprintf(os.Stderr, "\"host:port\"\n")
    fmt.Fprintf(os.Stderr, "Usage: %s \"utf16Client\" ", os.Args[0])
    fmt.Fprintf(os.Stderr, "\"host:port\"\n")
    fmt.Fprintf(os.Stderr, "Usage: %s \"jsonEchoClient|gobEchoClient\" ", os.Args[0])
    fmt.Fprintf(os.Stderr, "\"host:port\"\n")

    fmt.Fprintf(os.Stderr, "Usage: %s \"lookupport\" \"network-type\" ", os.Args[0])
    fmt.Fprintf(os.Stderr, "\"service\"\n")
}
