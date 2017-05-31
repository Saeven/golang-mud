package saevenx

import (
	"log"
)

type Race struct {
	Name          string
	Faction       string
	hitpointRegen int
	vitalityRegen int
}

var raceList = map[string]*Race{
	"demon": {
		Name:          "Demon",
		Faction:       "Hell",
		hitpointRegen: 10,
		vitalityRegen: 5,

	},
}

func getRace(raceName string) *Race {
	if val, ok := raceList[raceName]; ok {
		return val
	}
	log.Fatal( "A non-existant race key was requested")
	return nil
}
