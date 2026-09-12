package main

import (
	"encoding/json"
	"fmt"
	"go-task/internal/project"
	"net/http"
)

func main() {
	repo := project.NewMemoryRepository()

	http.HandleFunc("/health", healthHandler)

	http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		projectHandler(w, r, repo)
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
		fmt.Println("decode error:", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
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
