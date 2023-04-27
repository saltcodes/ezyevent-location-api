package handler

import (
	"ezyevent-location-api/internal/model"
	"ezyevent-location-api/internal/util"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

var events = db.Collection("events")

func ListEvents(c *fiber.Ctx) error {
	var eventList []model.Event
	cursor, err := events.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(ctx, &eventList); err != nil {

		return util.CreateResponseMessage(c, 500, "internal server error", nil)
	}

	return util.CreateResponseMessage(c, 200, "success", eventList)
}

func CreateEvent(c *fiber.Ctx) error {
	eventType := new(model.Event)

	if err := c.BodyParser(eventType); err != nil {
		return util.CreateResponseMessage(c, 500, "internal server error", err.Error())
	}

	result, err := events.InsertOne(ctx, eventType)
	if err != nil {
		return util.CreateResponseMessage(c, 500, "internal server error", err.Error())
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		eventType.Id = oid.Hex()
	}

	return util.CreateResponseMessage(c, 200, "success", eventType)
}
