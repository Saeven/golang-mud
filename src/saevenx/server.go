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
	fmt.Println("New connection (%d)", server.ConnectionCount())
	return &newConnection
}

func (server *Server) onClientConnectionClosed(connection *Connection, err error) {
	// delete the connection
	fmt.Println( "There are this many items in the list", server.ConnectionCount())
	for i, conn := range server.connectionList {
		if conn == connection {
			fmt.Println( "Found dead connection at position ", i )
			server.connectionList[i] = server.connectionList[len(server.connectionList)-1]
			server.connectionList = server.connectionList[:len(server.connectionList)-1]
			fmt.Println("Dead connection ", server.ConnectionCount())
			break

		}
	}

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
