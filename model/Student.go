package model

type Student struct {
	Id           int    `json:"_id,omitempty"  bson:"_id,omitempty"`
	FirstName    string `json:"first_name,omitempty"  bson:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"  bson:"last_name,omitempty"`
	Faculty      string `json:"faculty,omitempty"  bson:"faculty,omitempty"`
	Year         int    `json:"year,omitempty"  bson:"year,omitempty"`
	CourseEnroll []int  `json:"course_enroll,omitempty"  bson:"course_enroll,omitempty"`
}
