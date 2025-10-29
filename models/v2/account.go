package v2

import (
	"time"

	"github.com/SatisfactoryServerManager/ssmcloud-resources/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountSchema struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	AccountName string             `json:"accountName" bson:"accountName"`
	JoinCode    string             `json:"joinCode" bson:"joinCode"`

	AgentIds       primitive.A                `json:"-" bson:"agents" mson:"collection=agents"`
	Agents         []AgentSchema              `json:"agents" bson:"-"`
	AuditIds       primitive.A                `json:"-" bson:"audit" mson:"collection=accountaudits"`
	Audits         []AccountAuditSchema       `json:"audit" bson:"-"`
	IntegrationIds primitive.A                `json:"-" bson:"integrations" mson:"collection=accountintegrations"`
	Integrations   []AccountIntegrationSchema `json:"integrations" bson:"-"`

	InactivityState AccountInactivityState `json:"inactivityState" bson:"inactivityState"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type AuditType string

const (
	AuditType_UserAddedToAccount            AuditType = "added-user"
	AuditType_UserRemovedFromAccount        AuditType = "removed-user"
	AuditType_IntegrationAddedToAccount     AuditType = "added-integration"
	AuditType_IntegrationRemovedFromAccount AuditType = "removed-integration"
	AuditType_AgentAddedToAccount           AuditType = "added-agent"
	AuditType_AgentRemoveFromAccount        AuditType = "removed-agent"
)

type AccountAuditSchema struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Type      AuditType          `json:"type" bson:"type"`
	Message   string             `json:"message" bson:"message"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}

type AccountIntegrationSchema struct {
	ID         primitive.ObjectID     `json:"_id" bson:"_id"`
	Name       string                 `json:"name" bson:"name"`
	Type       IntegrationType        `json:"type" bson:"type"`
	Url        string                 `json:"url" bson:"url"`
	EventTypes []IntegrationEventType `json:"eventTypes" bson:"eventTypes"`
	CreatedAt  time.Time              `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time              `json:"updatedAt" bson:"updatedAt"`
}

type AccountInactivityState struct {
	Inactive     bool      `json:"inactive" bson:"inactive"`
	DateInactive time.Time `json:"dateInactive" bson:"dateInactive"`
	DeleteDate   time.Time `json:"deleteDate" bson:"deleteDate"`
}

func NewAccount(accountName string) *AccountSchema {
	return &AccountSchema{
		ID:             primitive.NewObjectID(),
		AccountName:    accountName,
		JoinCode:       utils.RandStringBytes(16),
		AgentIds:       make(primitive.A, 0),
		AuditIds:       make(primitive.A, 0),
		IntegrationIds: make(primitive.A, 0),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}
