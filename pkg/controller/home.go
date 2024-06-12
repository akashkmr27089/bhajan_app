package controller

import (
	"bhajann/pkg/services"

	"encoding/json"
	"net/http"
)

type HomeController struct {
	CategoryService *services.CategoryService
	ContentService  *services.ContentService
}

func NewHomeController(
	categoryService *services.CategoryService,
	contentService *services.ContentService,
) *HomeController {
	return &HomeController{
		CategoryService: categoryService,
		ContentService:  contentService,
	}
}

func (entity HomeController) ListCategories(
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()

	data, err := entity.CategoryService.Find(
		ctx,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (entity HomeController) AddCategories(
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()

	data, err := entity.CategoryService.Add(
		ctx,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (entity HomeController) ListContent(
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()

	data, err := entity.ContentService.Find(
		ctx,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (entity HomeController) AddContent(
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()

	data, err := entity.ContentService.Add(
		ctx,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
