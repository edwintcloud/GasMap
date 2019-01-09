package models

// Location is our location struct to be used by other models
type Location struct {
	Lat  string `json:"lat" bson:"lat"`
	Lng  string `json:"lng" bson:"lng"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}
