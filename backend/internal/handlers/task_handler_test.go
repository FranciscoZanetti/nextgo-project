package handlers

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
)

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
