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

// User represents a user in the system
type CategoryModel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	AlbumArt string             `bson:"album_art"`
	State    domain.State       `bson:"state"`
}

type CategoryModelService struct {
	entity *mongo.Collection
}

func NewCategoryModelService(in *mongo.Collection) *CategoryModelService {
	return &CategoryModelService{
		entity: in,
	}
}

func (model *CategoryModelService) GetCollection() *mongo.Collection {
	model.entity = config.Client.Database("test").Collection("category")
	return model.entity
}

func (model *CategoryModelService) Find(
	ctx context.Context,
	pagingDTO domain.PagingPointer,
) ([]CategoryModel, error) {
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

	var results []CategoryModel
	for docs.Next(ctx) {
		var model CategoryModel
		err := docs.Decode(&model)
		if err != nil {
			return nil, err
		}
		results = append(results, model)
	}

	return results, nil
}

func (model *CategoryModelService) InserVal(
	ctx context.Context,
	categoryModel CategoryModel,
) (*primitive.ObjectID, error) {
	marshal, err := bson.Marshal(categoryModel)
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
