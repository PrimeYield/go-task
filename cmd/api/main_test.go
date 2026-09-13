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

func TestProjectHandler_GetByID(t *testing.T) {
	repo := project.NewMemoryRepository()

	getEmptyReq := httptest.NewRequest(
		http.MethodGet,
		"/projects/1",
		nil,
	)

	getEmptyReq.SetPathValue("id", "1")

	getEmptyRecorder := httptest.NewRecorder()

	projectIDHandler(getEmptyRecorder, getEmptyReq, repo)

	if getEmptyRecorder.Code != http.StatusNotFound {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			getEmptyRecorder.Code,
		)
	}

	body := `{"name":"GoTask Backend"}`

	createReq := httptest.NewRequest(
		http.MethodPost,
		"/projects",
		strings.NewReader(body),
	)

	createRecorder := httptest.NewRecorder()

	projectHandler(createRecorder, createReq, repo)

	GetByIDReq := httptest.NewRequest(
		http.MethodGet,
		"/projects/1",
		nil,
	)

	GetByIDReq.SetPathValue("id", "1")

	getByIDRecorder := httptest.NewRecorder()
	projectIDHandler(getByIDRecorder, GetByIDReq, repo)

	if getByIDRecorder.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			getByIDRecorder.Code,
		)
	}
	var got project.Project

	err := json.NewDecoder(getByIDRecorder.Body).Decode(&got)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if got.ID != 1 {
		t.Errorf("expected ID %d, got %d", 1, got.ID)
	}

	if got.Name != "GoTask Backend" {
		t.Errorf(
			"expected name %q, got %q",
			"GoTask Backend",
			got.Name,
		)
	}
}

func TestProjectHandler_Update(t *testing.T) {
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

func TestProjectHandler_Delete(t *testing.T) {
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
