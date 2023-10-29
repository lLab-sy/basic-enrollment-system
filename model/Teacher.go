package model

type Teacher struct {
	Id        int    `json:"_id,omitempty"  bson:"_id,omitempty"`
	FirstName string `json:"first_name,omitempty"  bson:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"  bson:"last_name,omitempty"`
	Faculty   string `json:"faculty,omitempty"  bson:"faculty,omitempty"`
	OwnCourse []int  `json:"own_course,omitempty"  bson:"own_course,omitempty"`
}
