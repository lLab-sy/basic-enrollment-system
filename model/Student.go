package model

type Student struct {
	Id           string `json:"_id,omitempty"  bson:"_id,omitempty"`
	StudentId    string `json:"studentId,omitempty"  bson:"studentId,omitempty"`
	FirstName    string `json:"firstName,omitempty"  bson:"firstName,omitempty"`
	LastName     string `json:"lastName,omitempty"  bson:"lastName,omitempty"`
	Faculty      string `json:"faculty,omitempty"  bson:"faculty,omitempty"`
	Year         int    `json:"year,omitempty"  bson:"year,omitempty"`
	CourseEnroll []int  `json:"courseEnroll,omitempty"  bson:"courseEnroll,omitempty"`
}
