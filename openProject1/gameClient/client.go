package main

import (
	"encoding/json"
	. "fmt"
	"net"
	gsc "openProject1/graphic_server_communication"
	"os"
	"strings"
)

func main() {
	sAddr, err := net.ResolveUDPAddr("udp4", ":38735")
	if err != nil {
		Printf("Err: %#v", err)
		os.Exit(-1)
	}
	Printf("sAddr: %#v\n", sAddr)
	cAddr, err := net.ResolveUDPAddr("udp4", ":54874")
	if err != nil {
		Printf("Err: %#v", err)
		os.Exit(-1)
	}
	Printf("cAddr: %#v\n", cAddr)
	cConn, err := net.DialUDP("udp4", cAddr, sAddr)
	if err != nil {
		Printf("Err: %#v", err)
		os.Exit(-1)
	}
	Printf("cConn: %#v\n", cConn)
	cConn.Write([]byte("{\"ID\": \"123ID456\", \"Class\": \"PUTEPUTEPUTEPUTEPUTEPUTE\", \"md5\":\"987MD5\"}"))
	cConn.Close()
	ln, err := net.ListenUDP("udp4", cAddr)
	if err != nil {
		Printf("Err: %#v", err)
		os.Exit(-1)
	}
	Printf("ln: %#v\n", ln)
	b := make([]byte, 1024)
	data := gsc.GSIC_data{}
	for {
		size, addr, err := ln.ReadFrom(b)
		if err != nil {
			Println("Error Read")
		}

		if size != 0 {
			Printf("remote : %#v\n", addr)
			dec := json.NewDecoder(strings.NewReader(string(b[:])))
			dec.Decode(&data)

			Printf("Received %d bytes saying: \"%s\"\n", size, b)
			Printf("And the struct now contains:\n\t %#v\n", data)
		}

	}
}
