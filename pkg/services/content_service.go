package services

import (
	"bhajann/pkg/domain"
	"bhajann/pkg/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContentService struct {
	ContentModelService *repository.ContentModelService
}

func NewContentService(
	in *repository.ContentModelService,
) *ContentService {
	return &ContentService{
		ContentModelService: in, // Assign the value
	}
}

func (entity *ContentService) Find(
	ctx context.Context,
) ([]domain.ContentDTO, error) {
	contentModelDTOS, err := entity.ContentModelService.Find(
		ctx,
		domain.PagingPointer{
			Limit: 0,
		},
	)
	if err != nil {
		return nil, err
	}

	response := make([]domain.ContentDTO, len(contentModelDTOS))
	for idx, val := range contentModelDTOS {
		response[idx] = domain.ContentDTO{
			ID:         val.ID,
			Name:       val.Name,
			AlbumArt:   val.AlbumArt,
			ContentUrl: val.ContentUrl,
			CategoryID: val.CategoryID,
		}
	}
	return response, nil
}

func (entity *ContentService) Add(
	ctx context.Context,
) (*primitive.ObjectID, error) {

	data := repository.ContentModel{
		Name:       "Music",
		AlbumArt:   "http://google.com",
		ContentUrl: "http://google.com/asdf",
		CategoryID: "66675bf9c15fcab068782ab8",
	}

	insertedID, err := entity.ContentModelService.Insert(
		ctx,
		data,
	)
	if err != nil {
		return nil, err
	}

	return insertedID, nil
}
