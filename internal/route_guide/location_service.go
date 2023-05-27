package route_guide

import (
	"ezyevent-location-api/internal/proto"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
)

var eventsCollection = db.Collection("events-location")

func (s *Server) LocationData(ctx context.Context, data *proto.LocationObject) (*emptypb.Empty, error) {
	_, err := eventsCollection.InsertOne(context.TODO(), data)
	return new(emptypb.Empty), err
}
