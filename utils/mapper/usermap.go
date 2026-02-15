package mapper

import (
	models "github.com/SatisfactoryServerManager/ssmcloud-resources/models/v2"
	pbModels "github.com/SatisfactoryServerManager/ssmcloud-resources/proto/generated/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapUserSchemaToProto(u *models.UserSchema) *pbModels.User {
	if u == nil {
		return nil
	}

	return &pbModels.User{
		Id:              objectIDToString(u.ID),
		ExternalId:      u.ExternalID,
		Email:           u.Email,
		Username:        u.Username,
		ProfileImageUrl: u.ProfileImageStr,

		ApiKeys: mapUserAPIKeys(u.APIKeys),

		LastActive: timestamppb.New(u.LastActive),
		CreatedAt:  timestamppb.New(u.CreatedAt),
		UpdatedAt:  timestamppb.New(u.UpdatedAt),
	}
}

func mapUserAPIKeys(keys []models.UserAPIKey) []*pbModels.UserAPIKey {
	if len(keys) == 0 {
		return nil
	}

	result := make([]*pbModels.UserAPIKey, 0, len(keys))

	for _, k := range keys {
		result = append(result, &pbModels.UserAPIKey{
			ShortKey: k.ShortKey,
			Key:      k.Key, // Enable ONLY if required
		})
	}

	return result
}

func MapAccountSchemaToProto(a *models.AccountSchema) *pbModels.Account {
	if a == nil {
		return nil
	}

	return &pbModels.Account{
		Id:              objectIDToString(a.ID),
		AccountName:     a.AccountName,
		JoinCode:        a.JoinCode,
		Audit:           mapAccountAudits(a.Audits),
		Integrations:    mapAccountIntegrations(a.Integrations),
		InactivityState: mapAccountInactivityState(a.InactivityState),
		CreatedAt:       timestamppb.New(a.CreatedAt),
		UpdatedAt:       timestamppb.New(a.UpdatedAt),
	}
}

func mapAccountAudits(audits []models.AccountAuditSchema) []*pbModels.AccountAudit {
	if len(audits) == 0 {
		return nil
	}

	out := make([]*pbModels.AccountAudit, 0, len(audits))
	for i := range audits {
		a := &audits[i]
		out = append(out, &pbModels.AccountAudit{
			Id:        objectIDToString(a.ID),
			Type:      string(a.Type),
			Message:   a.Message,
			CreatedAt: timestamppb.New(a.CreatedAt),
		})
	}

	return out
}

func mapAccountIntegrations(integrations []models.AccountIntegrationSchema) []*pbModels.AccountIntegration {
	if len(integrations) == 0 {
		return nil
	}

	out := make([]*pbModels.AccountIntegration, 0, len(integrations))
	for i := range integrations {
		it := &integrations[i]
		// convert event types to strings
		ev := make([]string, 0, len(it.EventTypes))
		for j := range it.EventTypes {
			ev = append(ev, string(it.EventTypes[j]))
		}

		out = append(out, &pbModels.AccountIntegration{
			Id:         objectIDToString(it.ID),
			Name:       it.Name,
			Type:       int32(it.Type),
			Url:        it.Url,
			EventTypes: ev,
			CreatedAt:  timestamppb.New(it.CreatedAt),
			UpdatedAt:  timestamppb.New(it.UpdatedAt),
		})
	}

	return out
}

func mapAccountInactivityState(s models.AccountInactivityState) *pbModels.AccountInactivityState {
	return &pbModels.AccountInactivityState{
		Inactive:     s.Inactive,
		DateInactive: timestamppb.New(s.DateInactive),
		DeleteDate:   timestamppb.New(s.DeleteDate),
	}
}
