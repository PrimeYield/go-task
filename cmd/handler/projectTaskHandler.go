package handler

import (
	"encoding/json"
	"go-task/internal/project"
	"go-task/internal/task"
	"net/http"
	"strconv"
	"strings"
)

// func ProjectTasksHandler(
// 	w http.ResponseWriter,
// 	r *http.Request,
// 	repo project.Repository,
// ) {
// 	switch r.Method {
// 	case http.MethodPost:
// 		// POST /projects/{projectID}/tasks
// 		createTask(w, r, repo)
// 	// GET /projects/{projectID}/tasks
// 	case http.MethodGet:
// 		getProjectAllTasks(w, r, repo)

// 		// case http.MethodPut:

// 		// case http.MethodDelete:
// 	default:
// 		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 	}
// }

// post: /projects/{id}/tasks
// body: {"title":"title name"}
func CreateTaskHandler(
	w http.ResponseWriter,
	r *http.Request,
	repo project.Repository,
) {
	projectIDStr := r.PathValue("id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	var req task.CreateTaskRequest

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: "invalid request body",
		})
		return
	}
	if strings.TrimSpace(req.Title) == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: `request "Title" cannot be empty`,
		})
		return
	}

	newTask, exist := repo.CreateTask(req.Title, int64(projectID))
	if !exist {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: `project is not found.`,
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func GetProjectAllTasksHandler(
	w http.ResponseWriter,
	r *http.Request,
	repo project.Repository,
) {
	projectIDStr := r.PathValue("id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	tasks, exist := repo.GetProjectAllTasks(int64(projectID))
	if !exist {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: `project is not found.`,
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

// GET /projects/{id}/tasks/{task_id}
func GetProjectSingleTaskHandler(
	w http.ResponseWriter,
	r *http.Request,
	repo project.Repository,
) {
	projectIDStr := r.PathValue("id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: "project ID request Error: " + err.Error(),
		})
		return
	}
	taskIDStr := r.PathValue("task_id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: "task ID request Error" + err.Error(),
		})
		return
	}
	responseTask, exist := repo.GetProjectSingleTask(int64(projectID), int64(taskID))
	if !exist {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: "task ID is not exist",
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseTask)
}

func UpdateTaskHandler(
	w http.ResponseWriter,
	r *http.Request,
	repo project.Repository,
) {

	projectIDStr := r.PathValue("id")
	// var numID int64
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: "project ID request Error" + err.Error(),
		})
		return
	}

	taskIDStr := r.PathValue("task_id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: "task ID request Error" + err.Error(),
		})
		return
	}

	var req task.Task

	err = json.NewDecoder(r.Body).Decode(&req) // title & status
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: "invalid request body",
		})
		return
	}

	if strings.TrimSpace(req.Title) == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: `request "Title" cannot be empty`,
		})
		return
	}

	if req.Status != "todo" && req.Status != "doing" && req.Status != "done" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: `task status must be "todo" || "doing" || "done".`,
		})
		return
	}

	// numTaskID := int64(taskID)
	var uploadTask task.Task
	uploadTask, exist := repo.UpdateTask(int64(projectID), int64(taskID), req.Title, req.Status)
	if !exist {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: `project or task is not found.`,
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(uploadTask)
}

func DeleteTaskHandler(
	w http.ResponseWriter,
	r *http.Request,
	repo project.Repository,
) {
	projectIDStr := r.PathValue("id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: "project ID request Error : " + err.Error(),
		})
		return
	}
	taskIDStr := r.PathValue("task_id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: "task ID request Error : " + err.Error(),
		})
		return
	}
	success := repo.DeleteTask(int64(projectID), int64(taskID))
	if !success {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: "Delete Task Fail : project or task is not exist.",
		})
		return
	}

}
