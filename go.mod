module github.com/pellared/signalfx-go-tracing-test

go 1.16

require (
	github.com/signalfx/signalfx-go-tracing/contrib/mongodb/mongo-go-driver/mongo v1.7.0
	go.mongodb.org/mongo-driver v1.5.0
)

replace (
	github.com/signalfx/signalfx-go-tracing => github.com/pellared/signalfx-go-tracing v1.100.1
	github.com/signalfx/signalfx-go-tracing/contrib/mongodb/mongo-go-driver/mongo => github.com/pellared/signalfx-go-tracing/contrib/mongodb/mongo-go-driver/mongo v1.100.1
)
