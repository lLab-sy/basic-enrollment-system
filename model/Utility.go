package model

type Couter struct {
	Id  string `json:"_id,omitempty" bson:"_id,omitempty"`
	Seq int    `json:"seq,omitempty" bson:"seq,omitempty"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
