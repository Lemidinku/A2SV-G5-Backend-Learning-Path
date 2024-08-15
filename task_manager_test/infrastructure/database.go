
package infrastructure



import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"

)

func ConnectToMongoDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}


func CreateCollection(client *mongo.Client, databaseName string, collectionName string) (*mongo.Collection, error) {
	collection := client.Database(databaseName).Collection(collectionName)
	return collection, nil
}