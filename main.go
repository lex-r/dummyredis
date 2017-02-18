package main

import (
	"log"
	"net"
	"os"
)

func main() {

	ln, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatal("Error while listening port")
	}

	server := &Server{
		backend:  NewMemoryBackend(),
		listener: ln,
		log:      log.New(os.Stdout, "", log.LstdFlags),
	}

	server.Run()
}
