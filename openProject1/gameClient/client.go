package main

import (
	"bufio"
	"encoding/json"
	. "fmt"
	"net"
	gsc "openProject1/graphic_server_communication"
	"os"
	"strings"
	"time"
)

func check_error(err error) {
	if err != nil {
		Printf("Err: %#v", err)
		os.Exit(-1)
	}
}

func readFromStdin(textFromStdin chan string) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	check_error(err)
	textFromStdin <- text
}

func readFromServer(textFromServer chan gsc.GSIC_data, ln *net.UDPConn) {
	b := make([]byte, 1024)
	data := gsc.GSIC_data{}

	size, _, err := ln.ReadFrom(b)
	check_error(err)

	if size != 0 {
		dec := json.NewDecoder(strings.NewReader(string(b[:])))
		dec.Decode(&data)
		textFromServer <- data
		// Printf("Received %d bytes saying: \"%s\"\n", size, b)
	}
}

func main() {
	sAddr, err := net.ResolveUDPAddr("udp4", ":38735")
	check_error(err)
	// Printf("sAddr: %#v\n", sAddr)

	cWriteAddr, err := net.ResolveUDPAddr("udp4", ":54874")
	check_error(err)
	// Printf("cWriteAddr: %#v\n", cWriteAddr)

	cReadAddr, err := net.ResolveUDPAddr("udp4", ":54875")
	check_error(err)
	// Printf("cReadAddr: %#v\n", cReadAddr)

	cConn, err := net.DialUDP("udp4", cWriteAddr, sAddr)
	check_error(err)
	// Printf("cConn: %#v\n", cConn)

	cConn.Write([]byte("{\"ID\": \"123ID456\", \"Class\": \"STELIO EST UNE VIEILLE PUTE DEGARNIE.\", \"md5\":\"987MD5\"}"))
	defer cConn.Close()

	ln, err := net.ListenUDP("udp4", cReadAddr)
	// check_error(err)

	Printf("ln: %#v\n", ln)

	textFromStdin := make(chan string)
	textFromServer := make(chan gsc.GSIC_data)
	go readFromStdin(textFromStdin)
	go readFromServer(textFromServer, ln)

Loop:
	for {
		// Printf("In for\n")
		select {
		case textStdin := <-textFromStdin:
			{
				if textStdin == "exit\n" {
					Println("I'm out.")
					break Loop
				}
				cConn.Write([]byte(textStdin))
				Printf("text: %s\n", textStdin)
				go readFromStdin(textFromStdin)
				Println("Enter your command:")
			}
		case textServer := <-textFromServer:
			{
				go readFromServer(textFromServer, ln)
				Printf("And the struct now contains:\n\t %#v\n", textServer)
			}
		case <-time.After(1 * time.Millisecond):
			{
			}
		}
	}
	close(textFromServer)
	close(textFromStdin)
}
