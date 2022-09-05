package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Netflix struct {
	Id     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Series string             `json:"series,omitempty" bson:"series,omitempty"`

	// !omitempty for boolean omits false value
	Watched bool `json:"watched,omitempty" bson:"watched"`
}
