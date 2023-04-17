package models

import "time"

type Post struct{
	ID 			*int		`json:"id"`
	Title 		string		`json:"title"`
	Body 		string		`json:"body"`
	Image 		string		`json:"image"`
	ImageUrl 	string		`json:"image_url"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"updated_at"`
}