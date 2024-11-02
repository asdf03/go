package main

import {
    "log"
    "net/http"

    "github.com/asdf03/go/internal/delivery/http"
    "github.com/asdf03/go/internal/repository"
    "github.com/asdf03/go/internal/usecase"
}

func main() {
    repo := repository.NewTodoRepository()
    useCase := usecase.NewTodoUsecase(repo)
    handler := http.NewTodoHandler(useCase)

    http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            handler.GetTodos(w, r)
        case http.MethodPost:
            handler.CreateTodo(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

