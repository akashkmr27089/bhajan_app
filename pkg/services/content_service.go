package services

import (
	"bhajann/internal"
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
	search string,
	pagingDTO domain.PagingPointer,
) ([]domain.ContentDTO, error) {
	filter := internal.ContentFilter{
		NameSearch: search,
	}
	contentModelDTOS, err := entity.ContentModelService.Find(
		ctx,
		filter,
		pagingDTO,
	)
	if err != nil {
		return nil, err
	}

	response := make([]domain.ContentDTO, len(contentModelDTOS))
	for idx, val := range contentModelDTOS {
		response[idx] = domain.ContentDTO{
			ID:          val.ID,
			Name:        val.Name,
			AlbumArt:    val.AlbumArtUrl,
			ContentUrl:  val.ContentUrl,
			CategoryID:  val.CategoryID,
			Artist:      val.Artist,
			Description: val.Description,
		}
	}
	return response, nil
}

func (entity *ContentService) Add(
	ctx context.Context,
	addContent internal.ContentDTO,
) (*primitive.ObjectID, error) {

	data := repository.ContentModel{
		Name:        addContent.Name,
		AlbumArtUrl: addContent.AlbumArtUrl,
		ContentUrl:  addContent.ContentUrl,
		CategoryID:  addContent.CategoryID,
		Artist:      addContent.Artist,
		Description: addContent.Description,
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
