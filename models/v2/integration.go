package v2

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

type IntegrationEventSchema struct {
	ID              primitive.ObjectID     `json:"_id" bson:"_id,omitempty"`
	Type            IntegrationType        `json:"type" bson:"type"`
	EventType       IntegrationEventType   `json:"eventType" bson:"eventType"`
	IntegrationId   primitive.ObjectID     `bson:"integrationId"`
	URL             string                 `bson:"url"`
	Payload         map[string]interface{} `bson:"payload"`
	Status          string                 `bson:"status"` // pending, processing, sent, failed
	Attempts        int                    `bson:"attempts"`
	LastError       string                 `bson:"last_error,omitempty"`
	NextAttemptAt   time.Time              `bson:"next_attempt_at"`
	ProcessingUntil *time.Time             `bson:"processing_until,omitempty"`
	ProcessingBy    string                 `bson:"processing_by,omitempty"`
	CreatedAt       time.Time              `bson:"created_at"`
	SentAt          *time.Time             `bson:"sent_at,omitempty"`
}
