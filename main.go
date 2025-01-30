package main

import (
	"database/sql"
	"github.com/CodyBrunson/kanbanproject/internal/database"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type apiConfig struct {
	env ENV
	db  *database.Queries
}

func main() {
	cfg := apiConfig{
		env: loadEnv(),
	}

	newdb, err := sql.Open("postgres", cfg.env.DBUrl)
	if err != nil {
		panic(err)
	}

	cfg.db = database.New(newdb)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(cfg.env.RootFilePath)))
	mux.HandleFunc("POST /api/tasks", cfg.handlerCreateTask)
	mux.HandleFunc("GET /api/tasks", cfg.handlerGetAllTasks)
	mux.HandleFunc("GET /api/tasks/{id}", cfg.handlerGetTaskByID)
	mux.HandleFunc("PUT /api/tasks/{id}", cfg.handlerUpdateTaskByID)
	mux.HandleFunc("DELETE /api/tasks/{id}", cfg.handlerDeleteTaskByID)

	srv := &http.Server{
		Addr:    ":" + cfg.env.Port,
		Handler: loggingMiddlewareHandler(mux),
	}

	log.Printf("Serving on: http://localhost:%s", cfg.env.Port)
	log.Fatal(srv.ListenAndServe())

}

func loggingMiddlewareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s\t%s\t%s\tPayload Size: %v", r.Method, r.URL, r.RemoteAddr, r.ContentLength)
		next.ServeHTTP(w, r)
	})
}
