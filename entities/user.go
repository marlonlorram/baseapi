package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Password []byte             `bson:"password" json:"-"`
}

type AuthResult struct {
	User  *User  `json:"user"`
	Token string `json:"token,omitempty"`
}
