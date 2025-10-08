package v1

import (
	"time"

	"github.com/mrhid6/go-mongoose/mongoose"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	ExternalID string             `json:"eid" bson:"eid"`
	Email      string             `json:"email" bson:"email"`
	Password   string             `json:"-" bson:"password"`

	IsAccountAdmin bool `json:"isAccountAdmin" bson:"isAccountAdmin"`
	Active         bool `json:"active" bson:"active"`

	TwoFAState User2FAState `json:"twoFAState" bson:"twoFAState"`

	ProfileImageURL string `json:"profileImageUrl" bson:"profileImageUrl"`

	APIKeys []UserAPIKey `json:"apiKeys" bson:"apiKeys"`

	InviteCode string `json:"inviteCode" bson:"inviteCode"`

	LastActive time.Time `json:"lastActive" bson:"lastActive"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" bson:"updatedAt"`
}

type User2FAState struct {
	TwoFASetup  bool   `json:"twoFASetup" bson:"twoFASetup"`
	TwoFASecret string `json:"-" bson:"twoFASecret"`
}

type UserAPIKey struct {
	Key      string `json:"-" bson:"key"`
	ShortKey string `json:"shortKey" bson:"shortKey"`
}

func (obj *Users) AtomicDelete() error {

	if _, err := mongoose.DeleteOne(bson.M{"_id": obj.ID}, Users{}); err != nil {
		return err
	}

	return nil
}
