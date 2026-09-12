package main

import (
	"encoding/json"
	"go-task/internal/project"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProjectHandler_Create(t *testing.T) {
	repo := project.NewMemoryRepository()

	body := `{"name":"GoTask Backend"}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/projects",
		strings.NewReader(body),
	)

	recorder := httptest.NewRecorder()

	projectHandler(recorder, req, repo)

	if recorder.Code != http.StatusCreated {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusCreated,
			recorder.Code,
		)
	}
	var got project.Project

	err := json.NewDecoder(recorder.Body).Decode(&got)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
}
