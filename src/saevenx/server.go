package saevenx

import (
	"net"
	"time"
	"fmt"
	"io/ioutil"
	"os"
)

type Server struct {
	connectionList []*Connection
	ticker         *time.Ticker
	Motd		string
}

func CreateServer() *Server {
	server := &Server{}
	server.connectionList = make([]*Connection, 0)
	pwd, _ := os.Getwd()

	fmt.Println(pwd)

	fmt.Println("[CONFIG] Pulling MOTD")
	bytes, _ := ioutil.ReadFile(pwd + "/resources/MOTD")
	server.Motd = string(bytes)

	return server
}

func (server *Server) AddConnection(connection net.Conn) *Connection {
	newConnection := Connection{conn: connection, timeConnected: time.Now(), server: server, state: STATE_WELCOME}
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
	server.ticker = time.NewTicker(time.Millisecond * 3000)

	go func() {
		for range server.ticker.C {
			for _, c := range server.connectionList {
				fmt.Printf("[TICK] Running update tick on player (%s) at state [%d]\n", c.username, c.state )
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
