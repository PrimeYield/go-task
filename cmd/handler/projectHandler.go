package handler

import (
	"encoding/json"
	"errors"
	"go-task/internal/response/project"
	"log"
	"net/http"
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
	svc ProjectService,
) {
	// 承接Client JSON
	var req project.CreateProjectRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: "invalid request body",
		})
		return
	}

	// Service Layer
	createdProject, err := svc.CreateProject(r.Context(), req.Name)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")

		if errors.Is(err, project.ErrInvalidProjectName) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(project.ErrorResponse{
				Error: "project name cannot be empty",
			})
			return
		}

		log.Printf("create project failed: %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: "internal server error",
		})
		return
	}

	// Response
	resp := project.ResponseProject{
		ID:        createdProject.ID,
		Name:      createdProject.Name,
		TaskCount: createdProject.TaskCount,
		CreatedAt: createdProject.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("encode project response failed: %v", err)
	}
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
