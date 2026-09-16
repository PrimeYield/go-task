package task

import "time"

type Task struct {
	TaskID    int64     `json:"task_id"`    //server create
	Title     string    `json:"title"`      //client create & update
	Status    string    `json:"status"`     //server create & client update
	CreatedAt time.Time `json:"created_at"` //server create
	// ProjectID int64     `json:"project_id"` //client
}

//Post : /projects/{id}/tasks/{task_id}
type CreateTaskRequest struct {
	Title string `json:"title"` //expend project content (.body)
	//status default todo...
}

type ErrorTaskResponse struct {
	Error string `json:"error"`
}

//Put : /projects/{id}/tasks/{task_id}
//       body { "title":"Task Name","status":"3 types"}
type UploadTaskRequest struct {
	Title  string `json:"title"`  //reName
	Status string `json:"status"` //reStatus
}
