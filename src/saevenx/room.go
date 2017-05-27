package saevenx

type Room struct {
	id          int
	title       string
	description string
	exits       map[string]*Room
}

/**
 * Load rooms from database
 * @TODO Actually load them from database...
 */
func loadRooms() map[int]*Room {
	rooms := make(map[int]*Room, 5);
	rooms[1] = &Room{
		id:    1,
		title: Colorize("{LThe center of the world"),
		description: Colorize(`{L   Shadows dance throughout this deadened ward of the towering coniferous{n
{Lforest. Dark and grey, the sun's pallid light filters through in translucent{n
{Lrays that stream down to the needle-laden soil - the sole reminder of life{n
{Lin this dream-like stasis where at this level, in every chosen view, line but{n
{Lthousands of massive aged trunks for miles here on outwards - save for a dark{n
{Lcitadel shaped structure that lay northwards in the fray of this hallucinatory{n
{Lsetting.  As though protected by some illusory ward, it seems barely visible{n
{Lalthough its structure is still discernible - an obvious measure of the power{n
{Lthat lay therein.{n
{L   A large, {nstr{congl{n{Cy cle{n{cft cir{ncle{L lay before what seems to be the entrance to{n
{Lthe edifice, and through the low resonant wing; one could swear that is it{n
{Lactually humming.{n`),
	}

	return rooms
}

func (room *Room) showTo(player *Player) {

	player.connection.Write(room.title + "\n\n")
	player.connection.Write(room.description)

}
