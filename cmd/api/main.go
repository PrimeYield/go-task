package main

import (
	"fmt"
	"go-task/cmd/handler"
	"go-task/internal/database"
	"go-task/internal/response/project"
	"log"
	"net/http"
)

// POST /projects/{projectID}/tasks

func main() {
	db, err := database.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	respPrject := project.NewProjectResponse()
	respTask := task.NewTaskResponse()

	mux := http.NewServeMux()

	{
		mux.HandleFunc("GET /health", healthHandler)
	}
	
	{
		mux.HandleFunc("POST /projects", func(w http.ResponseWriter, r *http.Request) {
			handler.CreateProjectHandler(w, r, respPrject)
		})
		mux.HandleFunc("GET /projects", func(w http.ResponseWriter, r *http.Request) {
			handler.ListProjectsHandler(w, r, respPrject)
		})
		mux.HandleFunc("GET /projects/{id}", func(w http.ResponseWriter, r *http.Request) {
			handler.SearchProjectHandler(w, r, respPrject)
		})
		mux.HandleFunc("PUT /projects/{id}", func(w http.ResponseWriter, r *http.Request) {
			handler.UpdateProjectHandler(w, r, respPrject)
		})
		mux.HandleFunc("DELETE /projects/{id}", func(w http.ResponseWriter, r *http.Request) {
			handler.DeleteProjectHandler(w, r, respPrject)
		})
	}
	{
		mux.HandleFunc("POST /projects/{id}/tasks", func(w http.ResponseWriter, r *http.Request) {
			handler.CreateTaskHandler(w, r, respPrject)
		})
		mux.HandleFunc("GET /projects/{id}/tasks", func(w http.ResponseWriter, r *http.Request) {
			handler.GetProjectAllTasksHandler(w, r, respPrject)
		})
		mux.HandleFunc("GET /projects/{id}/tasks/{task_id}", func(w http.ResponseWriter, r *http.Request) {
			handler.GetProjectSingleTaskHandler(w, r, respPrject)
		})
		mux.HandleFunc("PUT /projects/{id}/tasks/{task_id}", func(w http.ResponseWriter, r *http.Request) {
			handler.UpdateTaskHandler(w, r, respPrject)
		})
		mux.HandleFunc("DELETE /projects/{id}/tasks/{task_id}", func(w http.ResponseWriter, r *http.Request) {
			handler.DeleteTaskHandler(w, r, respPrject)
		})
	}

	// http.HandleFunc("/health", healthHandler)

	// http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
	// 	handler.ProjectHandler(w, r, repo)
	// })

	// http.HandleFunc("/projects/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	handler.ProjectIDHandler(w, r, repo)
	// })

	// http.HandleFunc("/projects/{id}/tasks", func(w http.ResponseWriter, r *http.Request) {
	// 	handler.ProjectTasksHandler(w, r, repo) //createTask() , updateTask() , deleteTask() , searchTask()
	// })

	fmt.Println("GoTask API starting on :8080")

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
