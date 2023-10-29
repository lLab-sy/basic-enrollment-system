package model

type ClassTime struct {
	BeginClass string `json:"begin_class,omitempty"  bson:"begin_class,omitempty"`
	EndClass   string `json:"end_class,omitempty"  bson:"end_class,omitempty"`
	DayClass   string `json:"day_class,omitempty"  bson:"day_class,omitempty"`
	Classroom  string `json:"classroom,omitempty"  bson:"classroom,omitempty"`
}

type Section struct {
	Teacher       int         `json:"teacher,omitempty"  bson:"teacher,omitempty"`
	ClassInWeek   []ClassTime `json:"class_in_week,omitempty"  bson:"class_in_week,omitempty"`
	TotalCapacity int         `json:"total_capacity,omitempty"  bson:"total_capacity,omitempty"`
	Students      []int       `json:"students,omitempty"  bson:"students,omitempty"`
	Description   string      `json:"description,omitempty"  bson:"description,omitempty"`
}

type Course struct {
	Id            int       `json:"_id,omitempty"  bson:"_id,omitempty"`
	Name          string    `json:"name,omitempty"  bson:"name,omitempty"`
	Teachers      []int     `json:"teachers,omitempty"  bson:"teachers,omitempty"`
	StudySystem   []string  `json:"study_system,omitempty"  bson:"study_system,omitempty"`
	GradingSystem string    `json:"grading_system,omitempty"  bson:"grading_system,omitempty"`
	Credit        float32   `json:"credit,omitempty"  bson:"credit,omitempty"`
	Faculty       string    `json:"faculty,omitempty"  bson:"faculty,omitempty"`
	Gened         bool      `json:"gened,omitempty"  bson:"gened,omitempty"`
	Sections      []Section `json:"sections,omitempty"  bson:"sections,omitempty"`
	Description   string    `json:"description,omitempty"  bson:"description,omitempty"`
}
