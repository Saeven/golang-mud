package saevenx

import "fmt"

type Player struct {
	Name        string
	CurrentRoom int
	connection  *Connection
	inventory   []*Item
}

func (player *Player) setConnection(connection *Connection) {
	player.connection = connection
}

func (player *Player) getCurrentRoom() int {
	return player.CurrentRoom
}

func (player *Player) do(verb string, arguments []string) {
	command, err := getCommand(verb)
	if err != nil {
		player.connection.Write(fmt.Sprint(err))
		return
	}

	command.closure(player, arguments)
}
