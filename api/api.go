package api

import (
	"GolangBookApi/auth"
	"GolangBookApi/handler"
	"github.com/go-chi/chi"
)

func GetNewRoutes() *chi.Mux {
	return chi.NewRouter()
}

func RoutesAddress(router *chi.Mux) {

	router.Post("/login", auth.LoginHandler)

	router.Group(func(r chi.Router) {

		r.Use(auth.VerifyJWT)
		r.Post("/api/v1/books", handler.CreateBook)
		r.Get("/api/v1/books/{id}", handler.GetBook)
		r.Get("/api/v1/books", handler.ListOfBooks)
		r.Put("/api/v1/books/{id}", handler.UpdateBook)
		r.Delete("/api/v1/books/{id}", handler.DeleteBook)
	})
}
