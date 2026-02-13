package mapper

import (
	models "github.com/SatisfactoryServerManager/ssmcloud-resources/models/v2"
	pb "github.com/SatisfactoryServerManager/ssmcloud-resources/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapUserSchemaToProto(u *models.UserSchema) *pb.User {
	if u == nil {
		return nil
	}

	return &pb.User{
		Id:              objectIDToString(u.ID),
		ExternalId:      u.ExternalID,
		Email:           u.Email,
		Username:        u.Username,
		ProfileImageUrl: u.ProfileImageStr,

		ApiKeys:        mapUserAPIKeys(u.APIKeys),
		LinkedAccounts: mapAccounts(u.LinkedAccounts),
		ActiveAccount:  MapAccountSchemaToProto(&u.ActiveAccount),

		LastActive: timestamppb.New(u.LastActive),
		CreatedAt:  timestamppb.New(u.CreatedAt),
		UpdatedAt:  timestamppb.New(u.UpdatedAt),
	}
}

func mapUserAPIKeys(keys []models.UserAPIKey) []*pb.UserAPIKey {
	if len(keys) == 0 {
		return nil
	}

	result := make([]*pb.UserAPIKey, 0, len(keys))

	for _, k := range keys {
		result = append(result, &pb.UserAPIKey{
			ShortKey: k.ShortKey,
			Key:      k.Key, // Enable ONLY if required
		})
	}

	return result
}

func mapAccounts(accounts []models.AccountSchema) []*pb.Account {
	if len(accounts) == 0 {
		return nil
	}

	result := make([]*pb.Account, 0, len(accounts))

	for i := range accounts {
		result = append(result, MapAccountSchemaToProto(&accounts[i]))
	}

	return result
}

func MapAccountSchemaToProto(a *models.AccountSchema) *pb.Account {
	if a == nil {
		return nil
	}

	return &pb.Account{
		Id:          objectIDToString(a.ID),
		AccountName: a.AccountName,
		// Map other required fields
	}
}
