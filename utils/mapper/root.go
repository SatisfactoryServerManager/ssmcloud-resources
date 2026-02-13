package mapper

import "go.mongodb.org/mongo-driver/bson/primitive"

func objectIDToString(id primitive.ObjectID) string {
	if id.IsZero() {
		return ""
	}
	return id.Hex()
}
