module github.com/pellared/signalfx-go-tracing-test

go 1.16

require (
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/signalfx/golib v2.5.1+incompatible // indirect
	github.com/signalfx/signalfx-go-tracing/contrib/mongodb/mongo-go-driver/mongo v1.7.0
	github.com/tinylib/msgp v1.1.5 // indirect
	go.mongodb.org/mongo-driver v1.5.0
	golang.org/x/sys v0.0.0-20210309074719-68d13333faf2 // indirect
)

replace (
	github.com/signalfx/signalfx-go-tracing => github.com/pellared/signalfx-go-tracing v1.100.1
	github.com/signalfx/signalfx-go-tracing/contrib/mongodb/mongo-go-driver/mongo => github.com/pellared/signalfx-go-tracing/contrib/mongodb/mongo-go-driver/mongo v1.100.1
	gopkg.in/Shopify/sarama.v1 => github.com/Shopify/sarama v1.26.1
)
