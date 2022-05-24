package libs

import (
	"bufio"
	"log"
	"net"
)

func Echo2(conn net.Conn) {
	defer conn.Close()
	 reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err != nil {
	log.Fatalln("Unable to read data")
	}
	log.Printf("Read %d bytes: %s", len(s), s)
	log.Println("Writing data")
	 writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil {
	log.Fatalln("Unable to write data")
	}
	 writer.Flush()
	}
	/*
	// Bind to TCP port 20080 on all interfaces.
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for {
		// Wait for connection. Create net.Conn on connection established.
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go libs.Echo2(conn)
	}
	*/