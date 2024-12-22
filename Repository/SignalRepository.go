package Repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"omeggleClone-signallingServer/Exceptions"
	"omeggleClone-signallingServer/databases"
	"omeggleClone-signallingServer/models"
)

func FindARandomConnectionExcept(currentConnectionID string) (*string, *Exceptions.ApplicationException) {
	cursor, err := databases.Collection.Aggregate(context.TODO(), mongo.Pipeline{
		{{
			"$match", bson.M{"connectionId": bson.D{{"$ne", currentConnectionID}}},
		}},
		{{"$sample", bson.D{
			{"size", 1},
		}}},
	})
	if err != nil {
		return nil, Exceptions.NewApplicationException(500, "Unknown Error", &err)
	}
	defer cursor.Close(context.TODO())
	var availableConnection *models.Connections
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&availableConnection)
		if err != nil {
			return nil, Exceptions.NewApplicationException(500, "Unknown Error", &err)
		}
		break
	}
	if availableConnection == nil {
		return nil, Exceptions.NewApplicationException(404, "Not Connection Available", nil)
	}
	return &availableConnection.ConnectionId, nil
}

func CreateConnection(connectionID string) *Exceptions.ApplicationException {
	_, err := databases.Collection.InsertOne(context.TODO(), models.Connections{ConnectionId: connectionID})
	if err != nil {
		return Exceptions.NewApplicationException(500, "Unknown error", &err)
	}
	return nil
}
func DeactivateConnection(connectionID string) *Exceptions.ApplicationException {
	filter := bson.M{"connectionId": connectionID}

	_, err := databases.Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return Exceptions.NewApplicationException(500, "Unknown error", &err)
	}
	return nil
}
