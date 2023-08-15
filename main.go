package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/Tekitori19/hotel-reservation-BE/API"
	"github.com/Tekitori19/hotel-reservation-BE/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi= "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const userColl = "users"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))	
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	coll := client.Database(dbname).Collection(userColl)

	user := types.User{
		FirstName: "Dinh",
		LastName: "Dwcks",
	}
	res, err := coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(client, res)

	listenAddr := flag.String("listen", ":3000", "listen address")
	flag.Parse()

	app := fiber.New()
	api := app.Group("api/v1")

	api.Get("/user", API.HandleGetUsers)
	api.Get("/user/:id", API.HandleGetUser)
	app.Listen(*listenAddr)
}

