package saevenx

import (
	"fmt"
)

const (
	POS_PRONE    = 4
	POS_KNEELING = 3
	POS_RECLINED = 2
	POS_SITTING  = 1
	POS_STANDING = 0
)

type Player struct {
	Name         string
	CurrentRoom  int
	connection   *Connection
	inventory    []*Item
	race         *Race
	position     int
	hitPoints    int
	hitPointsMax int
	vitality     int
	vitalityMax  int
}

func (player *Player) setConnection(connection *Connection) {
	player.connection = connection
}

func (player *Player) getCurrentRoom() int {
	return player.CurrentRoom
}

/**
 * Execute a command on behalf of this player
 */
func (player *Player) do(verb string, arguments []string) {
	command, err := getCommand(verb)
	if err != nil {
		player.connection.Write(fmt.Sprint(err))
		return
	}

	command.closure(player, arguments)
}

func (player *Player) sendPrompt() {
	str := fmt.Sprintf("\n%s%s< %dh/%dH %dv/%dV Pos: %s%s%s >\n<>%s\n",
		FG_GREEN,
		MOD_FAINT,
		player.hitPoints,
		player.hitPointsMax,
		player.vitality,
		player.vitalityMax,
		getPositionString(player.position),
		FG_GREEN,
		MOD_FAINT,
		MOD_CLEAR,
	)
	player.connection.Write(str)

}

func (player *Player) adjustPosition(newPosition int) {
	var message string
	switch player.position {
	case POS_STANDING:
		switch newPosition {
		case POS_STANDING:
			message = "You're already standing."
		case POS_KNEELING:
			message = "You kneel down onto one leg."
		case POS_SITTING:
			message = "You settle into a sitting position."
		case POS_RECLINED:
			message = "You sit and recline onto the ground"
		case POS_PRONE:
			message = "You're knocked off your feet! The world is spinning..."
		}
	case POS_KNEELING:
		switch newPosition {
		case POS_STANDING:
			message = "You're back on your feet."
		case POS_KNEELING:
			message = "But? You're already kneeling!"
		case POS_SITTING:
			message = "You bend your other leg and drop into a seated position."
		case POS_RECLINED:
			message = "You roll sideways into a reclined position."
		case POS_PRONE:
			message = "You're thrown to the ground, your ears are ringing!"
		}
	case POS_SITTING:
		switch newPosition {
		case POS_STANDING:
			message = "You jump back onto your feet."
		case POS_KNEELING:
			message = "You push yourself forward onto one knee."
		case POS_SITTING:
			message = "You're already sitting!"
		case POS_RECLINED:
			message = "You hunker down backward, into a reclined position."
		case POS_PRONE:
			message = "The world flashes white as you're thrown backward!"
		}

	case POS_RECLINED:
		switch newPosition {
		case POS_STANDING:
			message = "With some effort, you're back on your feet!"
		case POS_KNEELING:
			message = "You roll about into a kneeling position."
		case POS_SITTING:
			message = "You pull yourself into a sitting position."
		case POS_RECLINED:
			message = "The genius behind that request is outstanding."
		case POS_PRONE:
			message = "You're violently throw into a daze. What happened!?!"
		}
	}
	player.position = newPosition
	player.connection.Write(message)
}

func getPositionString(position int) string {
	switch position {
	case POS_PRONE:
		return FG_RED + "prone" + MOD_CLEAR
	case POS_KNEELING:
		return FG_B_CYAN + "kneeling" + MOD_CLEAR
	case POS_RECLINED:
		return FG_B_CYAN + "reclined" + MOD_CLEAR
	case POS_SITTING:
		return FG_B_CYAN + "sitting" + MOD_CLEAR
	case POS_STANDING:
		return FG_CYAN + "standing" + MOD_CLEAR
	}
	return "Unknown"
}

func (player *Player) pulseUpdate() {
	fmt.Printf("This is for player %s", player.Name)
	if player.hitPoints < player.hitPointsMax {
		player.hitPoints = min(player.hitPoints+player.regenHP(), player.hitPointsMax)
		player.sendPrompt()
	}
}

func (player *Player) regenHP() int {
	return random_int(0, player.race.hitpointRegen)
}
