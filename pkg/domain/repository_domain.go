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
	ID       primitive.ObjectID `json:"id"`
	Name     string             `json:"name"`
	AlbumArt string             `json:"album_art"`
	State    State              `json:"state"`
}

type ContentDTO struct {
	ID         primitive.ObjectID `json:"_id"`
	Name       string             `json:"name"`
	AlbumArt   string             `json:"album_art"`
	ContentUrl string             `json:"content_url"`
	CategoryID string             `json:"cateogry_id"`
}
