package main

import (
	"encoding/json"
	. "fmt"
	"net"
	gsc "openProject1/graphic_server_communication"
	"os"
	"strconv"
	"strings"
	// "time"
)

func check_error(err error) {
	if err != nil {
		Printf("Err: %#v", err)
		os.Exit(-1)
	}
}

func getClientReturnAddr(addr net.Addr) net.Addr {
	port, err := strconv.ParseInt(addr.String()[strings.Index(addr.String(), ":")+1:], 10, 32)
	check_error(err)
	address := addr.String()[:strings.Index(addr.String(), ":")+1]
	addr, err = net.ResolveUDPAddr("udp4", Sprint(address, port+1))
	check_error(err)
	return addr
}

func treatIncomingComm(data gsc.GSIC_data, addr net.Addr, ln *net.UDPConn) {
	Printf("Handling communication #%s started\n", data.UniqueID)
	addr = getClientReturnAddr(addr)
	ln.WriteTo([]byte("{\"ID\": \"987ZYX321\", \"Class\": \"SALOPESALOPESALOPE\", \"md5\":\"123MD5\"}"), addr)
	Printf("Handling communication #%s ended\n", data.UniqueID)
}

func openCommLink(protocol, port string) *net.UDPConn {
	addr, err := net.ResolveUDPAddr(protocol, port)
	if err != nil {
		Printf("Err: %#v", err)
		os.Exit(-1)
	}
	Printf("addr: %#v\n", addr)
	ln, err := net.ListenUDP(protocol, addr)
	if err != nil {
		Printf("Err: %#v", err)
		os.Exit(-1)
	}
	Printf("ln: %#v\n", ln)
	return ln
}

func main() {
	Println("Server starting...")

	data := gsc.GSIC_data{}

	ln := openCommLink("udp4", ":38735")

	for {
		b := make([]byte, 1024) //we'll have to extend it
		for i := 0; i < 1024; i++ {
			b[i] = 0
		}
		size, addr, err := ln.ReadFrom(b)
		if err != nil {
			Println("Error Read")
		}

		if size != 0 {
			dec := json.NewDecoder(strings.NewReader(string(b[:])))
			dec.Decode(&data)

			Printf("Received %d bytes saying: \"%s\"\n", size, b)
			Printf("And the struct now contains:\n\t %#v\n", data)
			go treatIncomingComm(data, addr, ln)
		}
	}
}
