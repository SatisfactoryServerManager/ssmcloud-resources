package modelv2

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserSchema struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	ExternalID string             `json:"eid" bson:"eid"`
	Email      string             `json:"email" bson:"email"`
	Password   string             `json:"-" bson:"password"`

	ProfileImageURL string `json:"profileImageUrl" bson:"profileImageUrl"`

	APIKeys []UserAPIKey `json:"apiKeys" bson:"apiKeys"`

	LinkedAccountsIds primitive.A     `json:"-" bson:"linkedAccounts" mson:"collection=accounts"`
	LinkedAccounts    []AccountSchema `json:"linkedAccounts" bson:"-"`

	LastActive time.Time `json:"lastActive" bson:"lastActive"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" bson:"updatedAt"`
}

type UserAPIKey struct {
	Key      string `json:"-" bson:"key"`
	ShortKey string `json:"shortKey" bson:"shortKey"`
}
