package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/userAdityaa/code_quest_backend/handlers"
)

func InitializeRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", handlers.RootHandler)
	r.Get("/login/github/", handlers.GithubLoginHandler)
	r.Get("/login/github/callback", handlers.GithubCallBackHandler)
	return r
}
