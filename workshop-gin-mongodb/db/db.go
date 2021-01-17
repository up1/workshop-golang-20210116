package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

// Resource :: referene to database
type Resource struct {
	DB *mongo.Database
}

// Close :: to close database connection
func (r *Resource) Close() {
	logrus.Warning("Closing all db connections")
}

// CreateResource :: to create connection to database
func CreateResource() (*Resource, error) {
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	dbName := os.Getenv("MONGODB_DB_NAME")
	clusterEndpoint := os.Getenv("MONGODB_ENDPOINT")

	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)
	client, err := mongo.NewClient(
		options.Client().ApplyURI(connectionURI),
		options.Client().SetMinPoolSize(100),
		options.Client().SetMaxPoolSize(1000))
	if err != nil {
		logrus.Errorf("Failed to create client: %v", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		logrus.Errorf("Failed to connect to server: %v", err)
		return nil, err
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		logrus.Errorf("Failed to ping cluster: %v", err)
		return nil, err
	}

	return &Resource{DB: client.Database(dbName)}, nil
}
