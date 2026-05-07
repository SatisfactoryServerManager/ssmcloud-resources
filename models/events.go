package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type EventData struct {
	IntegrationId bson.ObjectID `json:"integrationId"`
	EventId       bson.ObjectID `json:"eventId"`
	EventType     string        `json:"eventType"`
	EventTime     time.Time     `json:"eventTime"`
}

type EventDataAgent struct {
	EventData
	AgentName string `json:"agentName"`
}
