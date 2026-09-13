package main

import (
	"encoding/json"
	"fmt"
	"go-task/internal/project"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	repo := project.NewMemoryRepository()

	http.HandleFunc("/health", healthHandler)

	http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		projectHandler(w, r, repo)
	})

	http.HandleFunc("/projects/{id}", func(w http.ResponseWriter, r *http.Request) {
		projectIDHandler(w, r, repo)
	})

	fmt.Println("GoTask API starting on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

func projectHandler(
	w http.ResponseWriter,
	r *http.Request,
	repo project.Repository,
) {
	switch r.Method {
	case http.MethodPost:
		createProject(w, r, repo)

	case http.MethodGet:
		listProjects(w, r, repo)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func createProject(
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

func listProjects(
	w http.ResponseWriter,
	r *http.Request,
	repo project.Repository,
) {
	projects := repo.List()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(projects)
}

func searchProject(
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

func uploadProject(
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

func deleteProject(
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
	exist := repo.Delete(int64(numID))
	if !exist {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(project.ErrorResponse{
			Error: `project is not found.`,
		})
		return
	}
	// w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func projectIDHandler(
	w http.ResponseWriter,
	r *http.Request,
	repo project.Repository,
) {
	switch r.Method {
	case http.MethodGet:
		searchProject(w, r, repo)

	case http.MethodPut:
		uploadProject(w, r, repo)
	case http.MethodDelete:
		deleteProject(w, r, repo)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
