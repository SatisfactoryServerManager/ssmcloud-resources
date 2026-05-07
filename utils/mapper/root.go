package mapper

import "go.mongodb.org/mongo-driver/v2/bson"

func objectIDToString(id bson.ObjectID) string {
	if id.IsZero() {
		return ""
	}
	return id.Hex()
}
