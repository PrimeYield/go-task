package project

type Project struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CreateProjectRequest struct {
	Name string `json:"name"`
}
