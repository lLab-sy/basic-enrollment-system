package model

type FacultyMember struct {
	FacultyName string    `json:"facultyName,omitempty"  bson:"facultyName,omitempty"`
	Teachers    []Teacher `json:"teachers,omitempty"  bson:"teachers,omitempty"`
	Students    []Student `json:"students,omitempty"  bson:"students,omitempty"`
}
