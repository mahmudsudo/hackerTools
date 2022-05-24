package main

import (
	"log"
	"net"
	"PortGo/libs"
)



func main() {
	 listener, err := net.Listen("tcp", ":80")
	  if err != nil {
log.Fatalln("Unable to bind to port")
}
for {
	conn, err := listener.Accept()
	if err != nil {
	log.Fatalln("Unable to accept connection")
	}
	go libs.Handle(conn)
	}
}
