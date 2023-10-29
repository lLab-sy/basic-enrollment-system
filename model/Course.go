package model

type ClassTime struct {
	BeginClass string `json:"beginClass,omitempty"  bson:"beginClass,omitempty"`
	EndClass   string `json:"endClass,omitempty"  bson:"endClass,omitempty"`
	DayClass   string `json:"dayClass,omitempty"  bson:"dayClass,omitempty"`
	Classroom  string `json:"classroom,omitempty"  bson:"classroom,omitempty"`
}

type Section struct {
	Teacher       int         `json:"teacher,omitempty"  bson:"teacher,omitempty"`
	ClassInWeek   []ClassTime `json:"classInWeek,omitempty"  bson:"classInWeek,omitempty"`
	TotalCapacity int         `json:"totalCapacity,omitempty"  bson:"totalCapacity,omitempty"`
	Students      []int       `json:"students,omitempty"  bson:"students,omitempty"`
	Description   string      `json:"description,omitempty"  bson:"description,omitempty"`
}

type Course struct {
	Id            string    `json:"_id,omitempty"  bson:"_id,omitempty"`
	CourseId      string    `json:"courseId,omitempty"  bson:"courseId,omitempty"`
	Name          string    `json:"name,omitempty"  bson:"name,omitempty"`
	Teachers      []int     `json:"teachers,omitempty"  bson:"teachers,omitempty"`
	StudySystem   []string  `json:"studySystem,omitempty"  bson:"studySystem,omitempty"`
	GradingSystem string    `json:"gradingSystem,omitempty"  bson:"gradingSystem,omitempty"`
	Credit        float32   `json:"credit,omitempty"  bson:"credit,omitempty"`
	Faculty       string    `json:"faculty,omitempty"  bson:"faculty,omitempty"`
	Gened         bool      `json:"gened,omitempty"  bson:"gened,omitempty"`
	Sections      []Section `json:"sections,omitempty"  bson:"sections,omitempty"`
	Description   string    `json:"description,omitempty"  bson:"description,omitempty"`
}
