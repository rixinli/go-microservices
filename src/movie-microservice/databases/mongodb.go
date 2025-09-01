/*
 * @File: databases.mongodb.go
 * @Description: Handles MongoDB connections
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package databases

import (
	"context"
	"time"

	"github.com/rixinli/go-microservices/src/movie-microservice/common"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB manages MongoDB connection
type MongoDB struct {
	Client *mongo.Client
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

	return err
}

// Close the existing connection
func (db *MongoDB) Close() {
	if db.Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
		defer cancel()
		_ = db.Client.Disconnect(ctx)
	}
}
