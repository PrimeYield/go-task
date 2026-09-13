package project

import "time"

type Project struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"create_at"`
}

type CreateProjectRequest struct {
	Name string `json:"name"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type UploadProjectRequest struct {
	Name string `json:"name"`
}
