package main

import (
	"encoding/json"
	"go-task/cmd/handler"
	"go-task/internal/project"
	"go-task/internal/task"
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

	handler.CreateProjectHandler(recorder, req, repo)

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

	handler.ListProjectsHandler(getEmptyRecorder, getEmptyReq, repo)

	if getEmptyRecorder.Code != http.StatusNotFound {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusNotFound,
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

	handler.CreateProjectHandler(createRecorder, createReq, repo)

	GetByIDReq := httptest.NewRequest(
		http.MethodGet,
		"/projects/1",
		nil,
	)

	GetByIDReq.SetPathValue("id", "1")

	getByIDRecorder := httptest.NewRecorder()
	handler.ListProjectsHandler(getByIDRecorder, GetByIDReq, repo)

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

	handler.CreateProjectHandler(recorder, req, repo)

	if recorder.Code != http.StatusCreated {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusCreated,
			recorder.Code,
		)
	}
	var createGot project.Project

	err := json.NewDecoder(recorder.Body).Decode(&createGot)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if createGot.Name != "GoTask Backend" {
		t.Errorf(
			"expected name %q, got %q",
			"GoTask Backend",
			createGot.Name,
		)
	}
	updateBody := `{"name":"IIoT Network"}`

	updateReq := httptest.NewRequest(
		http.MethodPut,
		"/projects",
		strings.NewReader(updateBody),
	)

	updateReq.SetPathValue("id", "1")

	updateRecorder := httptest.NewRecorder()

	handler.UpdateProjectHandler(updateRecorder, updateReq, repo)

	if updateRecorder.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			updateRecorder.Code,
		)
	}

	var updateGot project.Project
	err = json.NewDecoder(updateRecorder.Body).Decode(&updateGot)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if updateGot.Name != "IIoT Network" {
		t.Errorf(
			"expected name %q, got %q",
			"IIoT Network",
			updateGot.Name,
		)
	}
}

func TestProjectHandler_Delete(t *testing.T) {
	repo := project.NewMemoryRepository()

	{
		deleteBadReq := httptest.NewRequest(
			http.MethodDelete,
			"/projects/1",
			nil,
		)

		deleteBadReq.SetPathValue("id", "1")

		deleteBadRecorder := httptest.NewRecorder()

		handler.DeleteProjectHandler(deleteBadRecorder, deleteBadReq, repo)

		if deleteBadRecorder.Code != http.StatusNotFound {
			t.Fatalf(
				"expected status %d, got %d",
				http.StatusNotFound,
				deleteBadRecorder.Code,
			)
		}
	}

	body := `{"name":"GoTask Backend"}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/projects",
		strings.NewReader(body),
	)

	recorder := httptest.NewRecorder()

	handler.CreateProjectHandler(recorder, req, repo)

	if recorder.Code != http.StatusCreated {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusCreated,
			recorder.Code,
		)
	} //Create Project is success.

	deleteReq := httptest.NewRequest(
		http.MethodDelete,
		"/projects/1",
		nil,
	)

	deleteReq.SetPathValue("id", "1")

	deleteRecorder := httptest.NewRecorder()

	handler.DeleteProjectHandler(deleteRecorder, deleteReq, repo)

	if deleteRecorder.Code != http.StatusNoContent {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusNoContent,
			recorder.Code,
		)
	}

	var got project.Project

	err := json.NewDecoder(recorder.Body).Decode(&got)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
}

func TestTaskHandler_Create(t *testing.T) {
	repo := project.NewMemoryRepository()

	//create project
	{
		body := `{"name":"GoTask Backend"}`
		req := httptest.NewRequest(
			http.MethodPost,
			"/projects",
			strings.NewReader(body),
		)
		recorderProject := httptest.NewRecorder()
		handler.CreateProjectHandler(recorderProject, req, repo)

		if recorderProject.Code != http.StatusCreated {
			t.Fatalf(
				"expected status %d, got %d",
				http.StatusCreated,
				recorderProject.Code,
			)
		}
	}

	// create task in projects/1
	{
		bodyTask := `{"title":"Task A"}`
		reqTask := httptest.NewRequest(
			http.MethodPost,
			"/projects/1/tasks",
			strings.NewReader(bodyTask),
		)

		reqTask.SetPathValue("id", "1")

		recorderTask := httptest.NewRecorder()
		handler.CreateTaskHandler(recorderTask, reqTask, repo)

		if recorderTask.Code != http.StatusCreated {
			t.Fatalf(
				"expected status %d, got %d",
				http.StatusCreated,
				recorderTask.Code,
			)
		}
		var got task.Task

		err := json.NewDecoder(recorderTask.Body).Decode(&got)
		if err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if got.TaskID != 1 {
			t.Fatalf(
				"expected 1 task, got %d",
				got.TaskID,
			)
		}
		if got.Status != "todo" {
			t.Fatalf(
				"expected todo status, got %s",
				got.Status,
			)
		}
	}

	//create task in projects/2 not exist
	{
		bodyTask := `{"title":"Task A"}`
		reqTask := httptest.NewRequest(
			http.MethodPost,
			"/projects/2/tasks",
			strings.NewReader(bodyTask),
		)

		reqTask.SetPathValue("id", "2")

		recorderTask := httptest.NewRecorder()
		handler.CreateTaskHandler(recorderTask, reqTask, repo)

		if recorderTask.Code != http.StatusNotFound {
			t.Fatalf(
				"expected status %d, got %d",
				http.StatusNotFound,
				recorderTask.Code,
			)
		}
		var got project.Project

		err := json.NewDecoder(recorderTask.Body).Decode(&got)
		if err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}
	}
}

func TestTaskHandler_GetProjectAllTasks(t *testing.T) {
	repo := project.NewMemoryRepository()

	//create project
	{
		body := `{"name":"GoTask Backend"}`
		req := httptest.NewRequest(
			http.MethodPost,
			"/projects",
			strings.NewReader(body),
		)
		recorderProject := httptest.NewRecorder()
		handler.CreateProjectHandler(recorderProject, req, repo)

		if recorderProject.Code != http.StatusCreated {
			t.Fatalf(
				"expected status %d, got %d",
				http.StatusCreated,
				recorderProject.Code,
			)
		}
	}

	// create task in projects/1
	{
		bodyTask := `{"title":"Task A"}`
		reqTask := httptest.NewRequest(
			http.MethodPost,
			"/projects/1/tasks",
			strings.NewReader(bodyTask),
		)

		reqTask.SetPathValue("id", "1")

		recorderTask := httptest.NewRecorder()
		handler.CreateTaskHandler(recorderTask, reqTask, repo)

		if recorderTask.Code != http.StatusCreated {
			t.Fatalf(
				"expected status %d, got %d",
				http.StatusCreated,
				recorderTask.Code,
			)
		}
		var got task.Task

		err := json.NewDecoder(recorderTask.Body).Decode(&got)
		if err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if got.TaskID != 1 {
			t.Fatalf(
				"expected 1 task, got %d",
				got.TaskID,
			)
		}
		if got.Status != "todo" {
			t.Fatalf(
				"expected todo status, got %s",
				got.Status,
			)
		}
	}

	//create task in projects/2 not exist
	{
		bodyTask := `{"title":"Task A"}`
		reqTask := httptest.NewRequest(
			http.MethodPost,
			"/projects/2/tasks",
			strings.NewReader(bodyTask),
		)

		reqTask.SetPathValue("id", "2")

		recorderTask := httptest.NewRecorder()
		handler.CreateTaskHandler(recorderTask, reqTask, repo)

		if recorderTask.Code != http.StatusNotFound {
			t.Fatalf(
				"expected status %d, got %d",
				http.StatusNotFound,
				recorderTask.Code,
			)
		}
		var got project.Project

		err := json.NewDecoder(recorderTask.Body).Decode(&got)
		if err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}
	}
}
