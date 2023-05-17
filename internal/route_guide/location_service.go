package route_guide

import (
	"ezyevent-location-api/internal/proto"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

var eventsCollection = db.Collection("events-location")

func (s *Server) LocationData(ctx context.Context, data *proto.LocationObject) (*emptypb.Empty, error) {
	_, err := eventsCollection.InsertOne(context.TODO(), data)
	return new(emptypb.Empty), err
}

func (s *Server) FindEventsWithin(ctx context.Context, locationData *proto.LocationQueryObject) (proto.EventsLists, error) {
	var locationObject []*proto.LocationObject
	queryEventsWithin(&ctx, locationData, locationObject)
	return proto.EventsLists{
		Events: locationObject,
	}, nil
}

func queryEventsWithin(ctx *context.Context, queryObject *proto.LocationQueryObject, eventsList []*proto.LocationObject) {
	query := bson.D{
		{"location",
			bson.D{
				{"$near",
					bson.D{
						{"$geometry",
							bson.D{
								{"type", "Point"},
								{"coordinates",
									bson.A{
										queryObject.GetLatitude(),
										queryObject.GetLongitude(),
									},
								},
							},
						},
						{"$minDistance", 0},
						{"$maxDistance", queryObject.Radius},
					},
				},
			},
		},
	}

	cursor, err := eventsCollection.Find(*ctx, query)

	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(*ctx, eventsList); err != nil {
		log.Fatal(err)
	}
}
