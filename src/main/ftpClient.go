/* FTPClient
 */
package main

import (
    "fmt"
    "net"
    "os"
    "bufio"
    "strings"
    "bytes"
)

// strings used by the user interface
const (
    uiDir  = "dir"
    uiCd   = "cd"
    uiPwd  = "pwd"
    uiQuit = "quit"
    uiHelp = "help"
)

func ftpClient(host string) {

    conn, err := net.Dial("tcp", host+":1202")
    checkError(err)

    ftpClientHelp()

    reader := bufio.NewReader(os.Stdin)
    for {
        line, err := reader.ReadString('\n')
        // lose trailing whitespace
        line = strings.TrimRight(line, " \t\r\n")
        if err != nil {
            break
        }

        // split into command + arg
        strs := strings.SplitN(line, " ", 2)
        // decode user request
        switch strs[0] {
        case uiDir:
            dirRequest(conn)
        case uiCd:
            if len(strs) != 2 {
                fmt.Println("cd <dir>")
                continue
            }
            fmt.Println("CD \"", strs[1], "\"")
            cdRequest(conn, strs[1])
        case uiPwd:
            pwdRequest(conn)
        case uiQuit:
            conn.Close()
            os.Exit(0)
        case uiHelp:
            ftpClientHelp()
        default:
            fmt.Println("Unknown command")
        }
    }
}

func ftpClientHelp() {
    fmt.Println("Usage: dir|pwd|quit")
    fmt.Println("Usage: cd <dir>")
}

func dirRequest(conn net.Conn) {
    conn.Write([]byte(DIR + " "))

    var buf [512]byte
    result := bytes.NewBuffer(nil)
    for {
        // read till we hit a blank line
        n, _ := conn.Read(buf[0:])
        result.Write(buf[0:n])
        length := result.Len()
        contents := result.Bytes()
        if string(contents[length-4:]) == "\r\n\r\n" {
            fmt.Println(string(contents[0 : length-4]))
            return
        }
    }
}

func cdRequest(conn net.Conn, dir string) {
    conn.Write([]byte(CD + " " + dir))
    var response [512]byte
    n, _ := conn.Read(response[0:])
    s := string(response[0:n])
    if s != "OK" {
        fmt.Println("Failed to change dir")
    }
}

func pwdRequest(conn net.Conn) {
    conn.Write([]byte(PWD))
    var response [512]byte
    n, _ := conn.Read(response[0:])
    s := string(response[0:n])
    fmt.Println("Current dir \"" + s + "\"")
}
