package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	mudcore "github.com/Saeven/golang-mud/src/saevenx"
	"github.com/spf13/cobra"
)

var rootContext context.Context

func main() {
	rand.Seed(time.Now().Unix())

	rootContext = context.Background()
	rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use:   "golang-mud",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		service := ":7777"

		tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
		checkError(err)

		listener, err := net.ListenTCP("tcp", tcpAddr)
		checkError(err)
		defer listener.Close()

		mudcore.GetServer().Start()
		listenForConnections(listener)
	},
}

func listenForConnections(listener *net.TCPListener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		newDescriptor(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func newDescriptor(connection net.Conn) {
	mudcore.ServerInstance.AddConnection(connection)
}
