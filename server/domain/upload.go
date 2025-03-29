package domain

import "mime/multipart"

type UploadReq struct {
	Type string                `form:"type" binding:"required"`
	File *multipart.FileHeader `form:"file" binding:"required"`
}

type Upload struct {
	FileName string `json:"file_name"`
	FileURL  string `json:"file_url"`
}
