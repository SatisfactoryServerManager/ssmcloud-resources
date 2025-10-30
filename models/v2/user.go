package v2

import (
	"html/template"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserSchema struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	ExternalID string             `json:"eid" bson:"eid"`
	Email      string             `json:"email" bson:"email"`
	Username   string             `json:"username" bson:"username"`

	ProfileImageURL template.URL `bson:"-" json:"-"` // not stored directly
	ProfileImageStr string       `bson:"profileImageUrl" json:"profileImageUrl"`

	APIKeys []UserAPIKey `json:"apiKeys" bson:"apiKeys"`

	LinkedAccountIds primitive.A     `json:"-" bson:"linkedAccounts" mson:"collection=accounts"`
	LinkedAccounts   []AccountSchema `json:"linkedAccounts" bson:"-"`

	ActiveAccountId primitive.ObjectID `json:"-" bson:"activeAccount" mson:"collection=accounts"`
	ActiveAccount   AccountSchema      `json:"activeAccount" bson:"-"`

	LastActive time.Time `json:"lastActive" bson:"lastActive"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" bson:"updatedAt"`
}

type UserAPIKey struct {
	Key      string `json:"-" bson:"key"`
	ShortKey string `json:"shortKey" bson:"shortKey"`
}
