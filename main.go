package main

import (
	"log"

	"github.com/labstack/echo/v4"

	echoinstr "github.com/signalfx/signalfx-go-tracing/contrib/labstack/echo.v4"
	"github.com/signalfx/signalfx-go-tracing/ddtrace/tracer"
)

func main() {
	r := echo.New()

	// Use the tracer middleware with your desired service name.
	r.Use(echoinstr.Middleware(echoinstr.WithServiceName("my-web-app")))

	// Set up an endpoint.
	r.GET("/hello", func(c echo.Context) error {
		return c.String(200, "hello world!")
	})

	r.GET("/image/encode", func(c echo.Context) error {
		// create a child span to track an operation
		span, _ := tracer.StartSpanFromContext(c.Request().Context(), "image.encode")

		// encode an image ...

		// finish the child span
		span.Finish()

		return c.String(200, "ok!")
	})

	// ...and listen for incoming requests
	if err := r.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
