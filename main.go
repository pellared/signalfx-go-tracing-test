package main

import (
	"context"
	"log"
	"time"

	mongotrace "github.com/signalfx/signalfx-go-tracing/contrib/mongodb/mongo-go-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	opts := options.Client()
	opts.SetMonitor(mongotrace.NewMonitor())
	c, err := mongo.Connect(ctx, opts.ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	if _, err := c.ListDatabases(ctx, nil); err != nil {
		log.Fatal(err)
	}
}
