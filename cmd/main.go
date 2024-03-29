package main

import (
	"net/http"

	"example.com/internal/adapters/handler"
	"example.com/internal/adapters/repository"
	"example.com/internal/core/services"
	"github.com/go-chi/chi"
)

var (
	uvc *services.UserService
	rvc *services.RecordService
)

func main() {

	r := chi.NewRouter()
	store := repository.NewPostgresRepository()
	uvc = services.NewUserService(store)
	rvc = services.NewRecordService(store)
	UserHandler := handler.NewHTTPUserHandler(*uvc)
	RecordHandler := handler.NewHTTPRecordHandler(*rvc)

	r.Route("/user", func(r chi.Router) {
		r.Get("/{id}", UserHandler.GetUser)
		r.Get("/", UserHandler.GetUsers)
		r.Put("/", UserHandler.UpdateUser)
		r.Post("/", UserHandler.AddUser)
		r.Delete("/{id}", UserHandler.DeleteUser)
	})
	r.Route("/record", func(r chi.Router) {
		r.Post("/", RecordHandler.AddRecord)
	})

	http.ListenAndServe(":8080", r)
}
