package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type FileEntry struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	Username      string             `bson:"username"`
	FileName      string             `bson:"fileName"`
	File          string             `bson:"file"`
	CloudPublicID string             `bson:"cloudPublicId"`
}
