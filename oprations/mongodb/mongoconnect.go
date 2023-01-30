package mongodb

import (
	"SimpleTikTok/oprations/viperconfigread"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type autoIncrement struct {
	Name  string `bson:"name"`
	Value int64  `bson:"value"`
}

var MongoDBCollection *mongo.Collection

func init() {
	mongoConfig, err := viperconfigread.ConfigReadToMongoDB()
	if err != nil {
		logx.Errorf("[pkg]mongodb [func]init [msg]get MongoDB config, [err]%v", err)
	}
	url := fmt.Sprintf("mongodb://%v:%v@%v:%v", mongoConfig.MongoUserName, mongoConfig.MongoPwd, mongoConfig.MongoUrl, mongoConfig.MongoPort)
	mongoDB := mongoConfig.MongoDB
	mongoTable := mongoConfig.MongoTable
	MongoDBCollection, err = MongoDBConnect(mongoDB, mongoTable, url)
	if err != nil {
		logx.Errorf("[pkg]mongodb [func]init [msg]get MongoDB collection, [err]%v", err)
	}
}

func MongoDBConnect(database string, Table string, url string) (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logx.Errorf("[pkg]mongodb [func]MongoDBConnect [msg]connect mongodb failed, [err]%v", err)
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		logx.Errorf("[pkg]mongodb [func]MongoDBConnect [msg]ping mongodb failed, [err]%v", err)
		return nil, err
	}

	collection := client.Database(database).Collection(Table)
	filter := bson.D{{
		Key:   "name",
		Value: "auto_increment",
	}}
	num, _ := collection.CountDocuments(context.Background(), filter)
	if num == 0 {
		increment := autoIncrement{
			Name:  "auto_increment",
			Value: 0,
		}
		collection.InsertOne(context.Background(), increment)
	}
	return collection, nil
}
