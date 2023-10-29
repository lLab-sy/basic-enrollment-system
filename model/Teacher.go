package model

type Teacher struct {
	Id        string `json:"_id,omitempty"  bson:"_id,omitempty"`
	TeacherId string `json:"teacherId,omitempty"  bson:"teacherId,omitempty"`
	FirstName string `json:"firstName,omitempty"  bson:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"  bson:"lastName,omitempty"`
	Faculty   string `json:"faculty,omitempty"  bson:"faculty,omitempty"`
	OwnCourse []int  `json:"ownCourse,omitempty"  bson:"ownCourse,omitempty"`
}
