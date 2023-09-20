package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Member struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name,omitempty"`
	Karma int                `bson:"karma,omitempty"`
}
