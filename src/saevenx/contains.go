package saevenx

type Contains interface{
	getItems() []*Item
	hasSpace() int
	listContents() string
}