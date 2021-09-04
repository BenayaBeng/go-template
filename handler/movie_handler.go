package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rysmaadit/go-template/common/responder"
	"github.com/rysmaadit/go-template/model"
	"gorm.io/gorm"
)

func CreateMovie(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var movie model.Movie
		err := json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		db.Create(&movie)
		responder.NewHttpResponse(r, w, http.StatusCreated, &movie, nil)
	}
}

func GetMovie(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		var movie model.Movie
		db.Where(&model.Movie{Slug: params["slug"]}).First(&movie)
		responder.NewHttpResponse(r, w, http.StatusNotFound, params["slug"], nil)
	}
}

func UpdateMovie(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var movie model.Movie
		db.Model(&movie).Where(&model.Movie{Slug: params["slug"]}).Update(&movie)
		responder.NewHttpResponse(r, w, http.StatusCreated, params["slug"], nil)
	}
}

func DeleteMovie(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var movie model.Movie
		db.Where(&model.Movie{Slug: params["slug"]}).Delete(&movie)
		responder.NewHttpResponse(r, w, http.StatusOK, "success", nil)
	}
}
