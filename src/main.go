package main

import (
	"crypto/tls"
	"fmt"
	"math"
	"os"
	"time"
)

func tlsVersionName(v uint16) string {
	switch v {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TSL 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	}
	return "Unsupported"
}

func connect(host, port string, out chan tls.ConnectionState) {
	conn, err := tls.Dial("tcp", host+":"+port, nil)
	if err != nil {
		panic("failed to connect: " + err.Error())
	}
	defer conn.Close()
	out <- conn.ConnectionState()
	defer close(out)
}

func sslDump(state *tls.ConnectionState) {
	fmt.Printf("TLS version : %s\n", tlsVersionName(state.Version))
	fmt.Printf("Cipher Suite: %s\n", tls.CipherSuiteName(state.CipherSuite))
	fmt.Printf("OCSP Stapling Support: %v\n", state.OCSPResponse != nil)

	serverCerts := state.PeerCertificates
	fmt.Println("\nServer certificate chain: ")
	for index, cert := range serverCerts {
		expiresIn := cert.NotAfter.Sub(time.Now())
		fmt.Printf("[%d](ca=%v): %s [Expires in: %v days] \n", index, cert.IsCA, cert.Subject.String(), math.Ceil(expiresIn.Hours()/24))
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage: ssl_checker <hostname/ip> <port>")
		os.Exit(0)
	}
	fmt.Printf("\nConnecting to %s ...\n\n", args[0])
	connChan := make(chan tls.ConnectionState)
	go connect(args[0], args[1], connChan)
	state := <-connChan
	sslDump(&state)
}
