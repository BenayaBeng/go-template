package router

import (
	"net/http"
	"os"

	"github.com/rysmaadit/go-template/handler"
	"github.com/rysmaadit/go-template/service"
	"gorm.io/gorm"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(dependencies service.Dependencies, db *gorm.DB) http.Handler {
	r := mux.NewRouter()

	setAuthRouter(r, dependencies.AuthService)
	setMovieRouter(r, db)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	return loggedRouter
}

func setAuthRouter(router *mux.Router, dependencies service.AuthServiceInterface) {
	router.Methods(http.MethodGet).Path("/auth/token").Handler(handler.GetToken(dependencies))
	router.Methods(http.MethodPost).Path("/auth/token/validate").Handler(handler.ValidateToken(dependencies))
}

func setMovieRouter(router *mux.Router, db *gorm.DB) {
	router.Methods(http.MethodPost).Path("/movie").Handler(handler.CreateMovie(db))
	router.Methods(http.MethodPost).Path("/movie/slug").Handler(handler.GetMovie(db))
	// router.Methods(http.MethodPatch, http.MethodPut).Path("/movie/{slug}").Handler(handler.UpdateMovie(dependencies))
	// router.Methods(http.MethodDelete).Path("/movie/{slug}").Handler(handler.DeleteMovie(dependencies))
}
