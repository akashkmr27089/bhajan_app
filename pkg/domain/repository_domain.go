package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type State string

const (
	ActiveState   State = "active"
	DisabledState State = "disable"
)

type CategoryDTO struct {
	ID       primitive.ObjectID
	Name     string
	AlbumArt string
	State    State
}

type ContentDTO struct {
	ID          primitive.ObjectID
	Name        string
	AlbumArt    string
	ContentUrl  string
	CategoryID  string
	Artist      string
	Description string
}
