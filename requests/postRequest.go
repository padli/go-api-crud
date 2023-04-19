package requests

import "mime/multipart"

type PostRequest struct {
	Title    string                `form:"title" binding:"required"`
	Slug     string                `form:"slug" binding:"required,slugUnique"`
	Body     string                `form:"body" binding:"required"`
	Image    *multipart.FileHeader `form:"image" binding:"required"`
	ImageUrl string                `form:"image_url"`
}
