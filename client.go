package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
)

func main() {
	var buffer = make([]byte, 1024)

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalf("client: dial: %s", err)
	}
	defer conn.Close()
	for {
		str := readInput()
		_, err := conn.Write([]byte(str))
		if err != nil {
			return
		}
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			return
		}
		response := string(buffer[0:bytesRead])
		fmt.Print(response)
		if response == "123\n" {
			log.Println("Encrypting connection")
			doEncrypted(conn)
		}
	}
}

func doEncrypted(unenc_conn net.Conn) {
	tc := tls.Config{InsecureSkipVerify: true}
	var buffer = make([]byte, 1024)
	conn := tls.Client(unenc_conn, &tc)
	err := conn.Handshake()
	if err != nil {
		log.Fatalf("tls: handshake: %s", err)
	}
	for {
		str := readInput()
		_, err := conn.Write([]byte(str))
		if err != nil {
			conn.Close()
			return
		}
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			conn.Close()
			return
		}
		response := string(buffer[0:bytesRead])
		fmt.Print(response)
		if response == "321\n" {
			log.Println("Decrypting connection")
			return
		}
	}
}

func readInput() string {
	var s string
	fmt.Scan(&s)
	return s
}
