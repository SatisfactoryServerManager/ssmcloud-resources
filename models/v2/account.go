package v2

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountSchema struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	AccountName string             `json:"accountName" bson:"accountName"`

	AgentIds       primitive.A                `json:"-" bson:"agents" mson:"collection=agents"`
	Agents         []AgentSchema              `json:"agents" bson:"-"`
	AuditIds       primitive.A                `json:"-" bson:"audit" mson:"collection=accountaudit"`
	Audits         []AccountAuditSchema       `json:"audit" bson:"-"`
	IntegrationIds primitive.A                `json:"-" bson:"integrations" mson:"collection=accountintegrations"`
	Integrations   []AccountIntegrationSchema `json:"integrations" bson:"-"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type AccountAuditSchema struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Type      string             `json:"type" bson:"type"`
	Message   string             `json:"message" bson:"message"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}

type IntegrationType int64

const (
	IntegrationWebhook IntegrationType = 0
	IntegrationDiscord IntegrationType = 1
)

type IntegrationEventType int64

const (
	IntegrationEventTypeAgentOnline  IntegrationEventType = 0
	IntegrationEventTypeAgentOffline IntegrationEventType = 1
	IntegrationEventTypePlayerJoined IntegrationEventType = 2
	IntegrationEventTypePlayerLeft   IntegrationEventType = 3
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
