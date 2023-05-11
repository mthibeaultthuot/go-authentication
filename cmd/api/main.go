package main

import (
	"github.com/D3xt3rrrr/go-cloud/authentication-server/data"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type Config struct {
	Client *mongo.Client
}

var (
	port = "0.0.0.0:9002"
)

func main() {
	app := Config{}

	server := &http.Server{
		Addr:    port,
		Handler: app.routes(),
	}

	log.Println("rest api server start on : ", port)

	mongodb := data.MongoDb{
		"MongoDb Tax",
		"mongodb://root:9144tbbw@mongodb:27017",
		"Tax",
		"user",
	}

	init, err := data.Init(mongodb)
	app.Client = init

	//go GrpcListen()

	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
