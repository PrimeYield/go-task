package handler

import (
	"encoding/json"
	"errors"
	"go-task/internal/response/project"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func SearchProjectHandler(
	w http.ResponseWriter,
	r *http.Request,
	svc ProjectService,
) {
	// 1. parse to int
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: "project ID must be a valid integer",
		})
		return
	}

	// 2. Service search
	foundProject, err := svc.GetProject(r.Context(), id)
	if err != nil {
		var status int
		var message string

		switch {
		case errors.Is(err, project.ErrInvalidProjectID):
			status = http.StatusBadRequest
			message = "project ID must be positive"

		case errors.Is(err, project.ErrProjectNotFound):
			status = http.StatusNotFound
			message = "project not found"

		default:
			log.Printf("get project failed: %v", err)
			status = http.StatusInternalServerError
			message = "internal server error"
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: message,
		})
		return
	}

	// 3. response format
	resp := project.ResponseProject{
		ID:        foundProject.ID,
		Name:      foundProject.Name,
		TaskCount: foundProject.TaskCount,
		CreatedAt: foundProject.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("encode project response failed: %v", err)
	}
}

func UpdateProjectHandler(
	w http.ResponseWriter,
	r *http.Request,
	repo project.Repository,
) {

	idText := r.PathValue("id")
	// var numID int64
	ID, err := strconv.Atoi(idText)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	var req project.UploadProjectRequest

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: "invalid request body",
		})
		return
	}

	if strings.TrimSpace(req.Name) == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: `request "Name" cannot be empty`,
		})
		return
	}

	numID := int64(ID)
	var uploadProject project.Project
	uploadProject, exist := repo.Update(numID, req.Name)
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
	json.NewEncoder(w).Encode(uploadProject)
}

func DeleteProjectHandler(
	w http.ResponseWriter,
	r *http.Request,
	repo project.Repository,
) {
	idText := r.PathValue("id")
	numID, err := strconv.Atoi(idText)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	// return errors.New("Project is not finished.")
	// return errors.New("Project is not exist.")
	err = repo.Delete(int64(numID))
	if err != nil {
		if errors.Is(err, project.ErrProjectNotFound) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(project.ErrorResponse{
				Error: err.Error(),
			})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	// w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

// func ProjectIDHandler(
// 	w http.ResponseWriter,
// 	r *http.Request,
// 	repo project.Repository,
// ) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		searchProject(w, r, repo)

// 	case http.MethodPut:
// 		uploadProject(w, r, repo)
// 	case http.MethodDelete:
// 		deleteProject(w, r, repo)
// 	default:
// 		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 	}
// }
