package models

import "time"

type Post struct {
	ID        *int      `json:"id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	Body      string    `json:"body"`
	Image     string    `json:"image"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MyTime time.Time

const (
	myTimeFormat = "Monday, 02 January 2006"
)

func (mt MyTime) MarshalJSON() ([]byte, error) {
	t := time.Time(mt)
	return []byte(`"` + t.Format(myTimeFormat) + `"`), nil
}
