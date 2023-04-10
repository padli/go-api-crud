package models

import (
	"time"
)

type Category struct{
	ID int
	Title string
	Desc string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}