package internal

import (
	"bhajann/pkg/domain"
	"context"
	"encoding/json"
	"io"
)

type ContentDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	AlbumArtUrl string `json:"album_art_url"`
	ContentUrl  string `json:"url"`
	Artist      string `json:"artist"`
	Description string `json:"description"`
	CategoryID  string `json:"cateogry_id"`
}

func (entity *ContentDTO) Populate(
	ctx context.Context,
	body io.ReadCloser,
) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(entity)
	if err != nil {
		return err
	}

	// todo: Yet to Implement
	// err = validate.Struct(entity)
	// if err != nil {
	// 	logger.Client.WithFields(logger.GetFailedEvent()).Error(err.Error())
	// 	return cerror.ValidationError(cerror.InvalidRequestMsg)
	// }
	return nil
}

type CategoryDTO struct {
	ID       string       `json:"id"`
	Name     string       `json:"name"`
	AlbumArt string       `json:"album_art"`
	State    domain.State `json:"state"`
}

func (entity *CategoryDTO) Populate(
	ctx context.Context,
	body io.ReadCloser,
) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(entity)
	if err != nil {
		return err
	}

	// todo: Yet to Implement
	// err = validate.Struct(entity)
	// if err != nil {
	// 	logger.Client.WithFields(logger.GetFailedEvent()).Error(err.Error())
	// 	return cerror.ValidationError(cerror.InvalidRequestMsg)
	// }
	return nil
}

type HomeScreenApiDTO struct {
	Category   []CategoryDTO `json:"category"`
	ContentDTO []ContentDTO  `json:"songs"`
}

type HomeScreenApiResponseDTO struct {
	Data *HomeScreenApiDTO `json:"data"`
}

func (entity *HomeScreenApiResponseDTO) ToDTO(
	categoryDTOS []domain.CategoryDTO,
	contentDTOS []domain.ContentDTO,
) {
	categoryDatas := make([]CategoryDTO, len(categoryDTOS))
	for idx, category := range categoryDTOS {
		categoryDatas[idx] = CategoryDTO{
			Name:     category.Name,
			AlbumArt: category.AlbumArt,
			State:    category.State,
			ID:       category.ID.Hex(),
		}
	}

	contentDatas := make([]ContentDTO, len(contentDTOS))
	for idx, content := range contentDTOS {
		contentDatas[idx] = ContentDTO{
			ID:          content.ID.Hex(),
			Name:        content.Name,
			AlbumArtUrl: content.AlbumArt,
			ContentUrl:  content.ContentUrl,
			CategoryID:  content.CategoryID,
			Artist:      content.Artist,
			Description: content.Description,
		}
	}
	entity.Data = &HomeScreenApiDTO{
		Category:   categoryDatas,
		ContentDTO: contentDatas,
	}
}
