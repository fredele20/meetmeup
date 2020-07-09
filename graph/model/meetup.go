package model

type Meetup struct {
	ID          string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      string `json:"userId"`
}
