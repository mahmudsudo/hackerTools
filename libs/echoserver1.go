package libs

import (
	"io"
	"log"
	"net"
)

func Echo(conn net.Conn) {
	defer conn.Close()

	b := make([]byte, 512)
	for {
		// Receive data via conn.Read into a buffer.
		size, err := conn.Read(b[0:])
		if err == io.EOF {
		log.Println("Client disconnected")
		break
		}
		if err != nil {
			log.Println("Unexpected error")
			break
			}
			log.Printf("Received %d bytes: %s \n", size, string(b))
			// Send data via conn.Write.
			log.Println("Writing data")
			if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
			}
		}
	}