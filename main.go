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

var (
	rootContext context.Context
	appVersion  string
	appName     string
)

const (
	defaultVersion = "0.0.0-dev"
	defaultAppName = "golang-mud"
)

func main() {
	if appVersion == "" {
		appVersion = defaultVersion
	}

	rand.Seed(time.Now().Unix())

	rootContext = context.Background()
	rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(versionCmd)
}

var rootCmd = &cobra.Command{
	Use:   appName,
	Short: "",
	Long:  "",
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run the mud server",
	Long:  "run the mud server",
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

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version information",
	Long:  "show version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s v%s\n", appName, appVersion)
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
