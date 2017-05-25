package saevenx

type Player struct {
	Name       string
	connection *Connection
}

func (player *Player) setConnection(connection *Connection) {
	player.connection = connection
}