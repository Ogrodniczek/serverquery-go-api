package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {

	//ip = "127.0.0.1"
	//port = 10011
	client_login_name = ''
	client_login_password = ''
	//sid = 1

	hostName := "127.0.0.1"
	portNum := "10011"

	command := "login" + client_login_name + client_login_password
	doDial(command, hostName, portNum)
}

func doDial(cmd, host, port string) {
	conn, err := net.Dial("tcp", host+":"+port)

	if err == nil {
		defer conn.Close()
		fmt.Fprintf(conn, cmd+"\r\n")
		scanner := bufio.NewScanner(conn)

		onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			if len(data) == 1 {
				return 0, nil, bufio.ErrFinalToken
			}
			if i := strings.Index(string(data), "\n\r"); i >= 0 {
				return i + 1, data[0:i], nil
			}
			return
		}
		scanner.Split(onComma)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

	} else {
		fmt.Printf("Some error %v", err)
		return

	}
}
