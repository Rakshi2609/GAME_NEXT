package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role string

const (
	RoleCreator  Role = "creator"
	RoleAttendee Role = "attendee"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"-"`
	Role     Role               `bson:"role" json:"role"`
}
