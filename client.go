package main

import (
	"fmt"
	"net"
	"strings"
)

// Client reads the commands from the client connection and execute them
type Client struct {
	connection net.Conn
	server     *Server
}

// Send sends data to client
func (c *Client) Send(v ...interface{}) {
	fmt.Fprintln(c.connection, v...)
}

// Handle handles the client connection. Here the main work is performed
func (c *Client) Handle() {
	var cmd, arg1, arg2 string
	defer c.connection.Close()

	for {
		arg1, arg2 = "", ""
		_, _ = fmt.Fscanf(c.connection, "%s", &cmd)
		fmt.Printf("Command: %s\n", cmd)

		switch strings.ToLower(cmd) {
		case "set":
			fmt.Fscanf(c.connection, "%s %s\n", &arg1, &arg2)
			c.server.backend.Set(&KVRecord{arg1, arg2})
		case "get":
			fmt.Fscanf(c.connection, "%s\n", &arg1)
			if record, ok := c.server.backend.Get(arg1); ok {
				c.Send(record.Value)
			} else {
				c.Send("Key not found")
			}
		case "keys":
			fmt.Fscanln(c.connection)
			records := c.server.backend.Keys()
			for _, record := range records {
				c.Send(record.Key)
			}
		case "quit":
			return
		}
	}
}
