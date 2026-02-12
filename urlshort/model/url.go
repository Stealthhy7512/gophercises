package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type URL struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	ShortURL  string        `bson:"short_url" json:"short_url" yaml:"short_url"`
	LongURL   string        `bson:"long_url" json:"long_url" yaml:"long_url"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at" yaml:"created_at"`
}
