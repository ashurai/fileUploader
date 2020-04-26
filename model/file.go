package model

import(
	"time"
)

type File struct {
	Name string
	Path string
	CreatedAt time.Time
}

type Files struct {
	ID interface{} `bson:"_id", json:"_id"`
	Name string
	Path string
	CreatedAt time.Time
}