package user

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID        bson.ObjectID `bson:"_id,omitempty"       json:"id"`
	Email     string        `bson:"email"               json:"email"`
	Password  string        `bson:"password,omitempty"  json:"-"`
	GoogleID  string        `bson:"google_id,omitempty" json:"-"`
	Username  string        `bson:"username"            json:"username"`
	Language  string        `bson:"language"            json:"language"`
	Theme     string        `bson:"theme"               json:"theme"`
	CreatedAt time.Time     `bson:"created_at"          json:"createdAt"`
	UpdatedAt time.Time     `bson:"updated_at"          json:"updatedAt"`
}

func New(email, username, hashedPassword string) *User {
	now := time.Now()
	return &User{
		ID:        bson.NewObjectID(),
		Email:     email,
		Password:  hashedPassword,
		Username:  username,
		Language:  "en",
		Theme:     "dark",
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func NewFromGoogle(email, username, googleID string) *User {
	now := time.Now()
	return &User{
		ID:        bson.NewObjectID(),
		Email:     email,
		GoogleID:  googleID,
		Username:  username,
		Language:  "en",
		Theme:     "dark",
		CreatedAt: now,
		UpdatedAt: now,
	}
}
