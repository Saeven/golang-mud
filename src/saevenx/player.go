package saevenx

import "fmt"

type Player struct {
	Name        string
	CurrentRoom int
	connection  *Connection
}

func (player *Player) setConnection(connection *Connection) {
	player.connection = connection
}

func (player *Player) getCurrentRoom() int {
	return player.CurrentRoom
}

func (player *Player) do(verb string, arguments []string) {

	command, error := getCommand(verb)
	if error != nil {
		player.connection.Write(fmt.Sprint(error))
		return
	}

	command.closure(player, arguments)
}
