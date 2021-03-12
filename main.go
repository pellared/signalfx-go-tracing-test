package main

import (
	"context"
	"log"

	mongotrace "github.com/signalfx/signalfx-go-tracing/contrib/mongodb/mongo-go-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	opts := options.Client()
	opts.SetMonitor(mongotrace.NewMonitor())
	if _, err := mongo.Connect(context.Background(), opts.ApplyURI("mongodb://localhost:27017")); err != nil {
		log.Fatal(err)
	}
}
