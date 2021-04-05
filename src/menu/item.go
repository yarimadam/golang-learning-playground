package menu

import "github.com/google/uuid"

type Item struct {
	ID       uuid.UUID
	Title    string
	Parent   *Item
	Children map[string]*Item
}