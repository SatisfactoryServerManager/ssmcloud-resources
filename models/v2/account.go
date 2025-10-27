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

type IntegrationType int64

const (
	IntegrationWebhook IntegrationType = 0
	IntegrationDiscord IntegrationType = 1
)

type IntegrationEventType string

const (
	IntegrationEventTypeAgentCreated IntegrationEventType = "ssm.agent.created"
	IntegrationEventTypeAgentRemoved IntegrationEventType = "ssm.agent.removed"
	IntegrationEventTypeAgentOnline  IntegrationEventType = "ssm.agent.online"
	IntegrationEventTypeAgentOffline IntegrationEventType = "ssm.agent.offline"
	IntegrationEventTypeUserAdded    IntegrationEventType = "ssm.user.added"
	IntegrationEventTypeUserRemoved  IntegrationEventType = "ssm.user.removed"
	IntegrationEventTypePlayerJoined IntegrationEventType = "game.player.joined"
	IntegrationEventTypePlayerLeft   IntegrationEventType = "game.player.left"
)

type AccountIntegrationSchema struct {
	ID         primitive.ObjectID        `json:"_id" bson:"_id"`
	Type       IntegrationType           `json:"type" bson:"type"`
	Url        string                    `json:"url" bson:"url"`
	EventTypes []IntegrationEventType    `json:"eventTypes" bson:"eventTypes"`
	Events     []AccountIntegrationEvent `json:"events" bson:"events"`
	CreatedAt  time.Time                 `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time                 `json:"updatedAt" bson:"updatedAt"`
}

type AccountIntegrationEvent struct {
	ID           primitive.ObjectID   `json:"_id" bson:"_id"`
	Type         IntegrationEventType `json:"type" bson:"type"`
	Retries      int                  `json:"retries" bson:"retries"`
	Status       string               `json:"status" bson:"status"`
	Data         interface{}          `json:"data" bson:"data"`
	ResponseData interface{}          `json:"responseData" bson:"responseData"`
	Completed    bool                 `json:"completed" bson:"completed"`
	Failed       bool                 `json:"failed" bson:"failed"`
	CreatedAt    time.Time            `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time            `json:"updatedAt" bson:"updatedAt"`
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
