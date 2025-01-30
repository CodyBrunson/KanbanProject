package main

import (
	"encoding/json"
	"github.com/CodyBrunson/kanbanproject/internal/database"
	"github.com/google/uuid"
	"net/http"
)

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

func (cfg *apiConfig) handlerGetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := cfg.db.GetAllTasks(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error getting tasks", err)
		return
	}

	if len(tasks) == 0 {
		respondWithJson(w, http.StatusNoContent, nil)
		return
	}
	var allTasks []Task
	for _, task := range tasks {
		allTasks = append(allTasks, Task{
			ID:          task.ID.String(),
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			CreatedAt:   task.CreatedAt.String(),
			UpdatedAt:   task.UpdatedAt.String(),
			DeletedAt:   task.DeletedAt.Time.String(),
		})
	}
	if len(allTasks) != 0 {
		respondWithJson(w, http.StatusOK, allTasks)
		return
	}
	respondWithJson(w, http.StatusNoContent, nil)
}

func (cfg *apiConfig) handlerGetTaskByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	parsedID, err := uuid.Parse(id)
	task, err := cfg.db.GetTaskByID(r.Context(), parsedID)
	if err != nil {
		respondWithError(w, http.StatusNoContent, "Task does not exist", err)
		return
	}

	respondWithJson(w, http.StatusOK, Task{
		ID:          task.ID.String(),
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt.String(),
		UpdatedAt:   task.UpdatedAt.String(),
		DeletedAt:   task.DeletedAt.Time.String(),
	})
}

func (cfg *apiConfig) handlerCreateTask(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Missing data in payload", err)
		return
	}

	newTask := database.CreateNewTaskParams{
		Title:       params.Title,
		Description: params.Description,
		Status:      "OPEN",
	}

	task, err := cfg.db.CreateNewTask(r.Context(), newTask)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating task", err)
	}

	respondWithJson(w, http.StatusCreated, Task{
		ID:          task.ID.String(),
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt.String(),
		UpdatedAt:   task.UpdatedAt.String(),
	})
}

func (cfg *apiConfig) handlerUpdateTaskByID(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	id := r.PathValue("id")
	parsedID, err := uuid.Parse(id)
	_, err = cfg.db.GetTaskByID(r.Context(), parsedID)
	if err != nil {
		respondWithError(w, http.StatusNoContent, "Task does not exist", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload", err)
	}

	updatedTask := database.UpdateTaskByIDParams{
		ID: parsedID,
	}
	if params.Title != "" {
		updatedTask.Title = params.Title
	}
	if params.Description != "" {
		updatedTask.Description = params.Description
	}

	err = cfg.db.UpdateTaskByID(r.Context(), updatedTask)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error updating task", err)
		return
	}

	respondWithJson(w, http.StatusOK, nil)
}

func (cfg *apiConfig) handlerDeleteTaskByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid task id", err)
		return
	}

	err = cfg.db.DeleteTask(r.Context(), parsedID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error deleting task", err)
		return
	}

	respondWithJson(w, http.StatusOK, nil)
}
