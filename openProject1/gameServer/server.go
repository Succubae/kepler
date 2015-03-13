package main

import (
	"encoding/json"
	"fmt"
	"net"
	gsc "openProject1/graphic_server_communication"
	em "openProject1/eventManager"
	"os"
	"strconv"
	"strings"
	"time"
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

func check_error(err error) {
	if err != nil {
		fmt.Printf("Err: %#v", err)
		os.Exit(-1)
	}
}

func getClientReturnAddr(addr net.Addr) net.Addr {
	port, err := strconv.ParseInt(addr.String()[strings.Index(addr.String(), ":")+1:], 10, 32)
	check_error(err)
	address := addr.String()[:strings.Index(addr.String(), ":")+1]
	addr, err = net.ResolveUDPAddr("udp4", fmt.Sprint(address, port+1))
	check_error(err)
	return addr
}

func treatIncomingComm(data gsc.GSIC_data, addr net.Addr, ln *net.UDPConn) {
	fmt.Printf("Handling communication #%s started\n", data.UniqueID)
	addr = getClientReturnAddr(addr)
	ln.WriteTo([]byte("{\"ID\": \"987ZYX321\", \"Class\": \"SALOPESALOPESALOPE\", \"md5\":\"123MD5\"}"), addr)
	fmt.Printf("Handling communication #%s ended\n", data.UniqueID)
}

func openCommLink(protocol, port string) *net.UDPConn {
	addr, err := net.ResolveUDPAddr(protocol, port)
	if err != nil {
		fmt.Printf("Err: %#v", err)
		os.Exit(-1)
	}
	fmt.Printf("addr: %#v\n", addr)
	ln, err := net.ListenUDP(protocol, addr)
	if err != nil {
		fmt.Printf("Err: %#v", err)
		os.Exit(-1)
	}
	fmt.Printf("ln: %#v\n", ln)
	return ln
}

func	receiveDataFromClients(b []byte, ln *net.UDPConn, data gsc.GSIC_data, received chan bool) {
	size, addr, err := ln.ReadFrom(b)
	if err != nil {
		fmt.Println("Error Read")
	}

	if size != 0 {
		dec := json.NewDecoder(strings.NewReader(string(b[:])))
		dec.Decode(&data)
		fmt.Printf("%s", color("red", "#############################\nReceived data !\n#############################\n"))
		fmt.Printf("%s", color("yellow", fmt.Sprintf("Received %d bytes saying: \"%s\"\n", size, b)))
		fmt.Printf("%s", color("green", fmt.Sprintf("And the struct now contains:\n\t %#v\n", data)))
		fmt.Printf("%s", color("red", "#############################\nEnd Communication!\n#############################\n"))
		go treatIncomingComm(data, addr, ln)
		received <- true
	}
}

func main() {
	fmt.Println("Server starting...")

	data := gsc.GSIC_data{}

	received := make(chan bool, 1)

//	ln := openCommLink("udp4", ":38735")
	ln := openCommLink("udp4", ":1500")

	b := make([]byte, 1024) //we'll have to extend it

	t := time.Now()

	mongo := em.ServInfo{"mongodb://user:pass@server.addr.com/db_name", nil}

	go receiveDataFromClients(b, ln, data, received)

	for {
		for i := 0; i < 1024; i++ {
			b[i] = 0
		}
		select {
		case <- received:
		{
			go receiveDataFromClients(b, ln, data, received)
		}
		case <-time.After(1 * time.Millisecond):
		{

		}
		}

		elapsed := time.Since(t)
		//fmt.Printf("elapsed time: %s\n", elapsed)
		if elapsed > time.Second * 5 {
			t = time.Now()
			go em.ParseGamesForEvents(false, mongo)
		}

	}
}
