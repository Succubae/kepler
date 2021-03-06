package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	gsc "openProject1/graphic_server_communication"
	"os"
	"strings"
	"time"
)

const (
	HOST = "localhost"
	PORT = "38735"
	TYPE = "tcp4"
)

func color(this_color, str string) string {
	switch this_color {
	case "yellow":
		return "\033[" + "33m" + str + "\033[0m"
	case "blue":
		return "\033[" + "34m" + str + "\033[0m"
	case "red":
		return "\033[" + "31m" + str + "\033[0m"
	case "green":
		return "\033[" + "32m" + str + "\033[0m"
	default:
		return str
	}
}

func	check_error(err error) {
	if err != nil {
		fmt.Printf("Err: %#v", err)
		os.Exit(-1)
	}
}

func	readFromStdin(textFromStdin chan string) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	check_error(err)
	if len(text) > 1 {
		textFromStdin <- text[:len(text) - 1]
	} else {
		textFromStdin <- text
	}
}

func	readFromServer(textFromServer chan gsc.GSIC_data, conn net.Conn) {
	b := make([]byte, 1024)
	data := gsc.GSIC_data{}

	size, err := conn.Read(b)
	check_error(err)

	if size != 0 {
		dec := json.NewDecoder(strings.NewReader(string(b[:])))
		dec.Decode(&data)
		fmt.Printf("%s", color("red", "#############################\nReceived data !\n#############################\n"))
		fmt.Printf("%s", color("yellow", fmt.Sprintf("Received %d bytes saying: \"%s\"\n", size, b)))
		fmt.Printf("%s", color("green", fmt.Sprintf("And the struct now contains:\n\t %#v\n", data)))
		fmt.Printf("%s", color("red", "#############################\nEnd Communication!\n#############################\n"))
		textFromServer <- data
		// Printf("Received %d bytes saying: \"%s\"\n", size, b)
	}
}

func	main() {
	conn, err := net.Dial(TYPE, HOST+":"+PORT)
	check_error(err)
	defer conn.Close()

	textFromServer := make(chan gsc.GSIC_data)
	textFromStdin := make(chan string)
	defer close(textFromServer)
	defer close(textFromStdin)
	go readFromStdin(textFromStdin)
	go readFromServer(textFromServer, conn)

Loop:
	for {
		select {
		case textStdin := <-textFromStdin:
		{
			if textStdin == "exit" {
				fmt.Println("I'm out.")
				break Loop
			}
			conn.Write([]byte(textStdin))
			fmt.Printf("text: %s\n", textStdin)
			go readFromStdin(textFromStdin)
			fmt.Println("Enter your command:")
		}
		case textServer := <-textFromServer:
		{
			go readFromServer(textFromServer, conn)
			fmt.Printf("And the struct now contains:\n\t %#v\n", textServer)
		}
		case <-time.After(1 * time.Millisecond):
		{
		}
		}
	}

}


//HIHIHI REMY EST UN GENIE !!!!
