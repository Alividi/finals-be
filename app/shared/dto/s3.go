package dto

import "io"

type FileUpload struct {
	Name        string
	Reader      io.Reader
	ContentType string
}
