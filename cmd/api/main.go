package main

import (
	"log"
	"net/http"

	"go-task/cmd/handler"
	"go-task/internal/database"
	"go-task/internal/service"
)

// 在編譯時確認兩個實作符合對應介面。
var _ service.ProjectStore = (*database.ProjectDB)(nil)
var _ handler.ProjectService = (*service.ProjectService)(nil)

func main() {
	// 1. open db
	db, err := database.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 2. pass to store
	store := database.NewProjectDB(db)

	// 3. pass to service
	svc := service.NewProjectService(store)

	// 4. regist router
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler)

	mux.HandleFunc("POST /projects", func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		handler.CreateProjectHandler(w, r, svc)
	})
	mux.HandleFunc("GET /projects/{id}", func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		handler.SearchProjectHandler(w, r, svc)
	})

	log.Println("Starting GoTask API on http://127.0.0.1:18080")

	if err := http.ListenAndServe("127.0.0.1:18080", mux); err != nil {
		log.Printf("HTTP server stopped: %v", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
