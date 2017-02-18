package main

import (
	"log"
	"net"
)

// Server accepts connections and provides the interaction with backend storage
type Server struct {
	backend  Backend
	listener net.Listener
	log      *log.Logger
}

// Run runs the server
func (s *Server) Run() {
	for {
		connection, err := s.listener.Accept()
		if err != nil {
			s.log.Printf("Error while accepting connection: %s", err)
			continue
		}

		s.log.Print("New connection accepted")

		client := &Client{
			connection: connection,
			server:     s,
		}

		go client.Handle()
	}
}
