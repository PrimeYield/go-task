package handler

import (
	"encoding/json"
	"errors"
	"go-task/internal/project"
	"net/http"
	"strconv"
	"strings"
)

func SearchProjectHandler(
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
	projectResponse, exist := repo.GetByID(int64(numID))
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
	json.NewEncoder(w).Encode(projectResponse)
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
