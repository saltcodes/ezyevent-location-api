package route_guide

import (
	"ezyevent-location-api/internal/database"
	"ezyevent-location-api/internal/proto"
)

var db = database.Config()

type Server struct {
	proto.UnimplementedLocationDataServiceServer
}
