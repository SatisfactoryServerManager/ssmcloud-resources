package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventData struct {
	IntegrationId primitive.ObjectID `json:"integrationId"`
	EventId       primitive.ObjectID `json:"eventId"`
	EventType     string             `json:"eventType"`
	EventTime     time.Time          `json:"eventTime"`
}

type EventDataAgent struct {
	EventData
	AgentName string `json:"agentName"`
}
