package main

import (
	"context"
	"log"
	"net/http"
	"time"

	mongotrace "github.com/signalfx/signalfx-go-tracing/contrib/mongodb/mongo-go-driver/mongo"
	httptrace "github.com/signalfx/signalfx-go-tracing/contrib/net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mux := httptrace.NewServeMux()
	mux.HandleFunc("/", handler)
	if err := http.ListenAndServe("localhost:8888", mux); err != nil {
		log.Println(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	opts := options.Client()
	opts.SetMonitor(mongotrace.NewMonitor())
	c, err := mongo.Connect(ctx, opts.ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Println(err)
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	if _, err := c.ListDatabases(ctx, nil); err != nil {
		log.Println(err)
	}
}
