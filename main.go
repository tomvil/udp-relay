package main

import (
	"flag"
	"fmt"
	"net"
)

var (
	localHost  = flag.String("localHost", "", "Local host to listen on")
	localPort  = flag.String("localPort", "", "Local port to listen on")
	remoteHost = flag.String("remoteHost", "", "Remote host to forward data to")
	remotePort = flag.String("remotePort", "", "Remote port to forward data to")

	bufferSize = flag.Int("bufferSize", 4096, "Buffer size for reading/writing data")

	debug = flag.Bool("debug", false, "Enable debug mode")
)

func udpRelay() {
	localAddr, err := net.ResolveUDPAddr("udp", *localHost+":"+*localPort)
	if err != nil {
		fmt.Println("Error resolving local address:", err)
		return
	}

	remoteAddr, err := net.ResolveUDPAddr("udp", *remoteHost+":"+*remotePort)
	if err != nil {
		fmt.Println("Error resolving remote address:", err)
		return
	}

	relayConn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		fmt.Println("Error listening on local address:", err)
		return
	}
	defer relayConn.Close()

	fmt.Printf("UDP relay listening on %s:%s and forwarding to %s:%s\n", *localHost, *localPort, *remoteHost, *remotePort)

	buffer := make([]byte, *bufferSize)

	for {
		n, clientAddr, err := relayConn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from local connection:", err)
			continue
		}

		_, err = relayConn.WriteToUDP(buffer[:n], remoteAddr)
		if err != nil {
			fmt.Println("Error forwarding data to remote host:", err)
			continue
		}

		if *debug {
			fmt.Printf("Relayed %d bytes from %s to %s\n", n, clientAddr, remoteAddr)
		}
	}
}

func main() {
	flag.Parse()

	if *localHost == "" || *localPort == "" || *remoteHost == "" || *remotePort == "" {
		fmt.Println("Usage: udp-relay -localHost <localHost> -localPort <localPort> -remoteHost <remoteHost> -remotePort <remotePort>")
		return
	}

	udpRelay()
}
