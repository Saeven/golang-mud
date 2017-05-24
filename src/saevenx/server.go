package saevenx

import (
	"net"
	"time"
	"fmt"
)

type Server struct {
	connectionList []*Connection
	ticker         *time.Ticker
}

func CreateServer() *Server {
	server := &Server{}
	server.connectionList = make([]*Connection, 0)
	return server
}

func (server *Server) AddConnection(connection net.Conn) *Connection {
	newConnection := Connection{conn: connection, timeConnected: time.Now(), server: server}
	server.connectionList = append(server.connectionList, &newConnection)
	go newConnection.listen()
	fmt.Printf("[CONN] There are %d connected users.\n", server.ConnectionCount())
	return &newConnection
}

func (server *Server) onClientConnectionClosed(connection *Connection, err error) {
	// delete the connection

	for i, conn := range server.connectionList {
		if conn == connection {
			server.connectionList[i] = server.connectionList[len(server.connectionList)-1]
			server.connectionList[len(server.connectionList)-1] = nil
			server.connectionList = server.connectionList[:len(server.connectionList)-1]
			break
		}
	}

	fmt.Printf( "[DISC] There are %d connected users.\n", server.ConnectionCount())

}

/**
 * Start the main game loop
 */
func (server *Server) Start() {
	server.ticker = time.NewTicker(time.Millisecond * 500)

	go func() {
		for range server.ticker.C {

			for _, c := range server.connectionList {
				c.Write(".")
			}
		}
	}()

}

func (server *Server) ConnectionCount() int {
	return len(server.connectionList)
}

/**
 * A message has been received on a given connection descriptor
 */
func (server *Server) onMessageReceived(connection *Connection, message string) {
	connection.Write("I received this string :" + message)
}
