package handler

import (
	"encoding/json"
	"go-task/internal/project"
	"net/http"
	"strings"
)

// func ProjectHandler(
// 	w http.ResponseWriter,
// 	r *http.Request,
// 	repo project.Repository,
// ) {
// 	switch r.Method {
// 	case http.MethodPost:
// 		CreateProjectHandler(w, r, repo)

// 	case http.MethodGet:
// 		ListProjectsHandler(w, r, repo)

// 	default:
// 		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 	}
// }

func CreateProjectHandler(
	w http.ResponseWriter,
	r *http.Request,
	repo project.Repository,
) {
	var req project.CreateProjectRequest

	err := json.NewDecoder(r.Body).Decode(&req)
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

	createdProject := repo.Create(req.Name)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(createdProject)
}

func ListProjectsHandler(
	w http.ResponseWriter,
	r *http.Request,
	repo project.Repository,
) {
	projects := repo.List()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(projects)
}
