package dto

import "mime/multipart"

type UserUpdateRequest struct {
	Id       int    `json:"-"`
	Email    string `form:"email" binding:"required,email"`
	FullName string `form:"full_name" binding:"required"`
	Phone    string `form:"phone" binding:"required"`
	Photo    multipart.File
}
