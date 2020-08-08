package models

type Measurement struct {
	ID          string `json:"id, omitempty" bson:"_id, omitempty"`
	Name        string `json:name, omitempty" bson:"name, omitempty"`
	Temperature string `json:temperature, omitempty" bson:"name, omitempty"`
	Humidity    string `json:humidity, omitempty" bson:"humidity, omitempty"`
}
