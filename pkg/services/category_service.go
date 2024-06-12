package services

import (
	"bhajann/internal"
	"bhajann/pkg/domain"
	"bhajann/pkg/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryService struct {
	CategoryServiceModel *repository.CategoryModelService
}

func NewCategoryService(in *repository.CategoryModelService) *CategoryService {
	return &CategoryService{
		CategoryServiceModel: in, // Assign the value
	}
}

func (entity *CategoryService) Find(
	ctx context.Context,
) ([]domain.CategoryDTO, error) {
	categoryModelDTOS, err := entity.CategoryServiceModel.Find(
		ctx,
		domain.PagingPointer{
			Limit: 0,
		},
	)
	if err != nil {
		return nil, err
	}

	response := make([]domain.CategoryDTO, len(categoryModelDTOS))
	for idx, val := range categoryModelDTOS {
		response[idx] = domain.CategoryDTO{
			ID:       val.ID,
			Name:     val.Name,
			AlbumArt: val.AlbumArt,
			State:    val.State,
		}
	}
	return response, nil
}

func (entity *CategoryService) Add(
	ctx context.Context,
	addCategorieDTO internal.CategoryDTO,
) (*primitive.ObjectID, error) {

	data := repository.CategoryModel{
		Name:     addCategorieDTO.Name,
		AlbumArt: addCategorieDTO.AlbumArt,
		State:    addCategorieDTO.State,
	}

	insertedID, err := entity.CategoryServiceModel.Insert(
		ctx,
		data,
	)
	if err != nil {
		return nil, err
	}

	return insertedID, nil
}
