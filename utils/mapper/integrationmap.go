package mapper

import (
	v2 "github.com/SatisfactoryServerManager/ssmcloud-resources/models/v2"
	pbModels "github.com/SatisfactoryServerManager/ssmcloud-resources/proto/generated/models"
	"github.com/SatisfactoryServerManager/ssmcloud-resources/utils"
)

func MapIntegrationEventsToProto(events []v2.IntegrationEventSchema) []*pbModels.IntegrationEvent {
	result := make([]*pbModels.IntegrationEvent, 0, len(events))
	for _, event := range events {
		result = append(result, MapIntegrationEventToProto(&event))
	}

	return result
}

func MapIntegrationEventToProto(integrationEvent *v2.IntegrationEventSchema) *pbModels.IntegrationEvent {

	return &pbModels.IntegrationEvent{
		Id:           integrationEvent.ID.Hex(),
		Type:         int32(integrationEvent.Type),
		EventType:    string(integrationEvent.EventType),
		Url:          integrationEvent.URL,
		Status:       integrationEvent.Status,
		Payload:      utils.ToJSON(integrationEvent.Payload),
		Response:     utils.ToJSON(integrationEvent.Response),
		ResponseCode: int32(integrationEvent.ResponseCode),
		Attempts:     int32(integrationEvent.Attempts),
		LastError:    integrationEvent.LastError,
	}
}
