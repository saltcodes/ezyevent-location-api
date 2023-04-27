package route_guide

import (
	"ezyevent-location-api/internal/proto"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) LocationData(ctx context.Context, data *proto.LocationObject) (*emptypb.Empty, error) {
	_, err := db.Collection("events-location").InsertOne(context.TODO(), data)
	return new(emptypb.Empty), err
}
