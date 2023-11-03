package route_guide

import (
	"ezyevent-location-api/internal/proto"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) FindEventsWithin(ctx context.Context, queryObject *proto.LocationQueryObject) (*proto.EventsLists, error) {
	locationObject := queryEventsWithin(&ctx, queryObject)
	return &proto.EventsLists{
		Events: locationObject,
	}, nil
}

func queryEventsWithin(ctx *context.Context, queryObject *proto.LocationQueryObject) []*proto.LocationObject {
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
						{"$maxDistance", queryObject.Radius * 1064},
					},
				},
			},
		},
	}

	cursor, err := eventsCollection.Find(*ctx, query)

	if err != nil {
		panic(err)
	}

	result := new([]*proto.LocationObject)
	if err = cursor.All(*ctx, result); err != nil {
		panic(err)
	}

	return *result
}

func (s *Server) DeleteEvent(ctx context.Context, eventID *proto.EventId) (*emptypb.Empty, error) {
	_, err := eventsCollection.DeleteOne(ctx, bson.D{{"id", eventID.Id}})
	return new(emptypb.Empty), err
}
