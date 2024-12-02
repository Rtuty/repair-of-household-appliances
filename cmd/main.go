package main

import (
	"modules/internal"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	localStorage := internal.NewLocalStorage()

	handlers := internal.GetHandlers(localStorage)

	// Регистрация обработчика для маршрута /application-form/create
	r.Post("/application/create", handlers.CreateApplicationForm)
	r.Get("/application/list", handlers.GetApplications)

	if err := http.ListenAndServe(":1323", r); err != nil {
		return
	}
}
