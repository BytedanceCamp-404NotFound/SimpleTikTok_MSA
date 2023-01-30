package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// get comment_id
func GetId(collection *mongo.Collection) (int64, error) {
	filter := bson.D{{
		Key:   "name",
		Value: "auto_increment",
	}}
	var temp autoIncrement
	err := collection.FindOne(context.Background(), filter).Decode(&temp)
	if err != nil {
		return -1, err
	}
	temp.Value++
	update := bson.D{{
		Key: "$inc",
		Value: bson.D{{
			Key:   "value",
			Value: 1,
		}},
	}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return -1, err
	}
	return temp.Value, nil
}
