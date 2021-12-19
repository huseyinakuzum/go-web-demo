package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "go-web-demo/docs"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	//Server
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, "/swagger/index.html")
	})

	//Init
	config := NewConfiguration()
	cbClient := NewCouchbaseClient(config.Couchbase)
	couchbaseRepository := NewCouchbaseRepositoryAdaptor(cbClient)
	reviewService := NewReviewService(couchbaseRepository)
	reviewController := NewReviewController(reviewService)
	reviewController.Register(e)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", config.ApplicationConfig.Server.Port),
	}

	if err := e.StartServer(srv); err != nil {
		panic(err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
