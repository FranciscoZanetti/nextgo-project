package handlers

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
	"os"
	"github.com/FranciscoZanetti/nextgo-project/backend/internal/database"
)

func TestMain(m *testing.M) {
	database.InitDB()

	database.DB.Exec("DELETE FROM task")

	database.DB.Exec(`
		INSERT INTO task (title, description) VALUES
		('Test Task 1', 'Description 1'),
		('Test Task 2', 'Description 2')
	`)

	code := m.Run()
	database.DB.Exec("DELETE FROM task")
	os.Exit(code)
}

func TestHandleTasks_Get(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
    w := httptest.NewRecorder()

    HandleTasks(w, req)

    resp := w.Result()
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusInternalServerError {
        t.Errorf("expected status 200 or 500, got %d", resp.StatusCode)
    }
}

func TestHandleTasks_Post(t *testing.T) {
    body := bytes.NewBufferString(`{"title":"New Task","description":"Details"}`)
    req := httptest.NewRequest(http.MethodPost, "/tasks", body)
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    HandleTasks(w, req)

    resp := w.Result()
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusInternalServerError {
        t.Errorf("expected status 200 or 500, got %d", resp.StatusCode)
    }
}

func TestHandleTaskByID_Get(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/tasks/1", nil)
    w := httptest.NewRecorder()

    HandleTaskByID(w, req)

    resp := w.Result()
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotFound {
        t.Errorf("expected status 200 or 404, got %d", resp.StatusCode)
    }
}

func TestHandleTaskByID_Put(t *testing.T) {
    body := bytes.NewBufferString(`{"title":"Updated","description":"Changed"}`)
    req := httptest.NewRequest(http.MethodPut, "/tasks/1", body)
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    HandleTaskByID(w, req)

    resp := w.Result()
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusInternalServerError {
        t.Errorf("expected status 204 or 500, got %d", resp.StatusCode)
    }
}

func TestHandleTaskByID_Delete(t *testing.T) {
    req := httptest.NewRequest(http.MethodDelete, "/tasks/1", nil)
    w := httptest.NewRecorder()

    HandleTaskByID(w, req)

    resp := w.Result()
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusInternalServerError {
        t.Errorf("expected status 204 or 500, got %d", resp.StatusCode)
    }
}
