package main

import (
	"encoding/json"
	. "fmt"
	"net"
	gsc "openProject1/graphic_server_communication"
	"os"
	"strings"
	"time"
)

func treatIncomingComm(data gsc.GSIC_data) {
	i := 0
	for i < 10 {
		Printf("Handling communication #%s\n", data.UniqueID)
		time.Sleep(1 * time.Second)
		i++
	}
	Printf("Handling done\n")
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

	b := make([]byte, 1024)
	data := gsc.GSIC_data{}

	ln := openCommLink("udp4", ":38735")

	for {
		size, err := ln.Read(b)
		if err != nil {
			Println("Error Read")
		}

		if size != 0 {
			remote := ln.RemoteAddr()
			net.
				Printf("remote : %#v\n", remote)
			dec := json.NewDecoder(strings.NewReader(string(b[:])))
			dec.Decode(&data)

			Printf("Received %d bytes saying: \"%s\"\n", size, b)
			Printf("And the struct now contains:\n\t %#v\n", data)
			go treatIncomingComm(data)
		}

	}
}
