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

func mapAccounts(accounts []models.AccountSchema) []*pbModels.Account {
	if len(accounts) == 0 {
		return nil
	}

	result := make([]*pbModels.Account, 0, len(accounts))

	for i := range accounts {
		result = append(result, MapAccountSchemaToProto(&accounts[i]))
	}

	return result
}

func MapAccountSchemaToProto(a *models.AccountSchema) *pbModels.Account {
	if a == nil {
		return nil
	}

	return &pbModels.Account{
		Id:          objectIDToString(a.ID),
		AccountName: a.AccountName,
		// Map other required fields
	}
}
