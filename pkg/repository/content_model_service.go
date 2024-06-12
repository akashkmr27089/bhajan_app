package repository

import (
	"bhajann/pkg/config"
	"bhajann/pkg/domain"
	"bhajann/pkg/repository/constants"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ContentModel struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	AlbumArt   string             `bson:"album_art"`
	ContentUrl string             `bson:"content_url"`
	CategoryID string             `bson:"cateogry_id"`
}

type ContentModelService struct {
	entity *mongo.Collection
}

func NewContentModelService(in *mongo.Collection) *ContentModelService {
	return &ContentModelService{
		entity: in,
	}
}

func (model *ContentModelService) GetDatabase() *mongo.Collection {
	return config.Client.Database("test").Collection("content")
}

func (model *ContentModelService) GetCollection() *mongo.Collection {
	model.entity = config.Client.Database("test").Collection("content")
	return model.entity
}

func (model *ContentModelService) Find(
	ctx context.Context,
	pagingDTO domain.PagingPointer,
) ([]ContentModel, error) {
	filterQuery := bson.M{}
	findOptions := options.Find()
	findOptions.SetMaxTime(1 * time.Second)

	if pagingDTO.LastID != nil {
		mongoId, err := primitive.ObjectIDFromHex(*pagingDTO.LastID)
		if err != nil {
			return nil, err
		}
		filterQuery[constants.AutoIDField] = bson.M{
			constants.MongoLtKeyword: mongoId,
		}
	}

	if pagingDTO.Limit > 0 {
		findOptions.SetLimit(int64(pagingDTO.Limit))
		findOptions.SetSort(
			bson.M{
				constants.AutoIDField: -1,
			},
		)
	}

	docs, err := model.entity.Find(
		ctx,
		filterQuery,
		findOptions,
	)

	if err != nil {
		return nil, err
	}

	var results []ContentModel
	for docs.Next(ctx) {
		var model ContentModel
		err := docs.Decode(&model)
		if err != nil {
			return nil, err
		}
		results = append(results, model)
	}

	return results, nil
}

func (model *ContentModelService) Insert(
	ctx context.Context,
	contentModelDTO ContentModel,
) (*primitive.ObjectID, error) {
	marshal, err := bson.Marshal(contentModelDTO)
	if err != nil {
		return nil, err
	}

	result, err := model.entity.InsertOne(
		ctx,
		marshal,
	)
	if err != nil {
		return nil, err
	}

	insertedID := result.InsertedID.(primitive.ObjectID)

	return &insertedID, nil
}
