package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Venue       string             `bson:"venue" json:"venue"`
	StartAt     time.Time          `bson:"startAt" json:"startAt"`
	PrizePool   string             `bson:"prizePool" json:"prizePool"`
	Logo        string             `bson:"logo" json:"logo"`
	Sponsors    []string           `bson:"sponsors" json:"sponsors"`
	Coordinator []string           `bson:"coordinator" json:"coordinator"`
	Speaker     string             `bson:"speaker" json:"speaker"`
	CreatedBy   primitive.ObjectID `bson:"createdBy" json:"createdBy"`
}
