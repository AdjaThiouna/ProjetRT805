package operation

import(
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Operation struct{
	Type string `bson:"type"`
	HasSucceeded bool `bson:"has_succeeded"`
}

type DeviceData struct{
	DeviceName string `bson:"device_name"`
	Operations []Operation `bson:"operations"`
}

func (d DeviceData) StoreToDatabase() error{
	client, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		return err
	}

	coll := client.Database("devices").Collection("operations")
	_, err = coll.InsertOne(context.Background(), d, nil)
	if err != nil {
		return err
	}
	return nil
}
