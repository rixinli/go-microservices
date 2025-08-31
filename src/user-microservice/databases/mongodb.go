/*
 * @File: databases.mongodb.go
 * @Description: Handles MongoDB connections
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package databases

import (
	"context"
	"time"

	"github.com/rixinli/go-microservices/src/user-microservice/common"
	"github.com/rixinli/go-microservices/src/user-microservice/models"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB manages MongoDB connection
type MongoDB struct {
	Client  *mongo.Client
	Databasename string
}

// Init initializes mongo database
func (db *MongoDB) Init() error {
	db.Databasename = common.Config.MgDbName

	// construct connection URI
	uri := "mongodb://" + common.Config.MgAddrs;
	if common.Config.MgDbUsername != "" && common.Config.MgDbPassword != "" {
        uri = "mongodb://" + common.Config.MgDbUsername + ":" + common.Config.MgDbPassword + "@" + common.Config.MgAddrs
    }

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
        log.Debug("Can't connect to mongo, go error: ", err)
        return err
    }
	
	// test connection
	if err := client.Ping(ctx, nil); err != nil {
        log.Debug("Mongo ping failed: ", err)
        return err
    }

    db.Client = client
    return db.initData()
}

// InitData initializes default data
func (db *MongoDB) initData() error{
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

	// query the the collection for Users
	collection := db.Client.Database(db.Databasename).Collection(common.ColUsers)

	// get the count of document of users
	count, err := collection.CountDocuments(ctx, bson.M{})
    if err != nil {
        return err
    }

	if count < 1 {
		user := models.User{
			ID:      primitive.NewObjectID(),
			Name: "admin",
			Password: "admin",
		}
		_,err = collection.InsertOne(ctx,user);
	}
	return err;
}

// Close the existing connection
func (db *MongoDB) Close() {
	if db.Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
		defer cancel()
		_ = db.Client.Disconnect(ctx)
	}
}
