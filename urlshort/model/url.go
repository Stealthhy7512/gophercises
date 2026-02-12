package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type URL struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"-"`
	ShortURL  string        `bson:"short_url"`
	LongURL   string        `bson:"long_url"`
	CreatedAt time.Time     `bson:"created_at"`
}
