package saevenx

type Room struct {
	id          int
	title       string
	description string
	exits       map[string]*Room
}