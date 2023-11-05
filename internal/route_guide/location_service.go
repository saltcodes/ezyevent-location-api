package route_guide

import (
	"ezyevent-location-api/internal/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
)

var eventsCollection = db.Collection("events-location")

func (s *Server) LocationData(ctx context.Context, data *proto.LocationObject) (*emptypb.Empty, error) {
	//filter for checking Id
	filter := bson.D{{"id", data.Id}}

	//Data prep
	d := bson.D{{"$set", data}}

	//Config for save or update
	opts := options.Update().SetUpsert(true)

	//Actual executions
	_, err := eventsCollection.UpdateOne(context.TODO(), filter, d, opts)
	return new(emptypb.Empty), err
}
