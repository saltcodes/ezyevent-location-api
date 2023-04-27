package model

type Location struct {
}

type Event struct {
	Id       string `bson:"_id,omitempty" json:"id"`
	Location struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"location"`
}
