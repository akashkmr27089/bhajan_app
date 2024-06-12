package controller

import (
	"bhajann/internal"
	"bhajann/pkg/domain"
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
	var addCategorieDTO internal.CategoryDTO
	err := addCategorieDTO.Populate(
		ctx,
		r.Body,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := entity.CategoryService.Add(
		ctx,
		addCategorieDTO,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (entity HomeController) ListContent(
	w http.ResponseWriter,
	request *http.Request,
) {
	ctx := request.Context()
	queryParams := request.URL.Query()
	pagingDTO := domain.GetPagingDTO(queryParams, 10)

	data, err := entity.ContentService.Find(
		ctx,
		pagingDTO,
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
	var addContentDTO internal.ContentDTO
	err := addContentDTO.Populate(
		ctx,
		r.Body,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := entity.ContentService.Add(
		ctx,
		addContentDTO,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (entity HomeController) HomePageApi(
	w http.ResponseWriter,
	request *http.Request,
) {
	ctx := request.Context()
	queryParams := request.URL.Query()
	pagingDTO := domain.GetPagingDTO(queryParams, 4)

	categoryDTOS, err := entity.CategoryService.Find(
		ctx,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	contentDTOS, err := entity.ContentService.Find(
		ctx,
		pagingDTO,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Combine Category and Content
	var response internal.HomeScreenApiResponseDTO
	response.ToDTO(categoryDTOS, contentDTOS)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
