package main

import (
        "bufio"
        "flag"
        "fmt"
        "net"
        "os"
        "strings"
)

func main() {
        host := flag.String("host", "127.0.0.1", "Host to connect to")
        port := flag.String("port", "4444", "Port to connect to")
        flag.Parse()
        server := *host + ":" + *port
        addr, err := net.ResolveTCPAddr("tcp", server)
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
        conn, err := net.DialTCP("tcp", nil, addr)
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
        prompt := "shell> "
        for {
                fmt.Printf("%s", prompt)
                reader := bufio.NewReader(os.Stdin)
                input, _ := reader.ReadString('\n')
                input = strings.Replace(input, "\n", "", -1)
                if input == "" {
                        continue
                } else if input == "exit" {
                        conn.Write([]byte("exit\n"))
                        os.Exit(0)
                }
                conn.Write([]byte(input + "\n"))
                recv(conn)
                println("")
        }
}

func recv(conn net.Conn) {
        reply := make([]byte, 1024)
        length, err := conn.Read(reply)
        if err != nil {
                println(err)
                os.Exit(1)
        }
        fmt.Print(string(reply))
        if length == 1024 {
                recv(conn)
        }
}
