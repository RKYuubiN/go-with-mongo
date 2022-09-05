package mongo

import (
	"context"
	"fmt"
	"log"

	"github.com/rkyuubin/gowithmongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	InsertOneSeries(ctx context.Context, movie model.Netflix) (*mongo.InsertOneResult, error)
	UpdateOneSeries(ctx context.Context, seriesId string) *mongo.UpdateResult
	DeleteOneSeries(ctx context.Context, seriesId string) *mongo.DeleteResult
	DeleteAllSeries(ctx context.Context) *mongo.DeleteResult
	GetAllSeries(ctx context.Context) []primitive.M
}

type repository struct {
	logger *log.Logger
}

func NewRepository(logger *log.Logger) Repository {
	return &repository{
		logger: logger,
	}
}

func (r *repository) InsertOneSeries(ctx context.Context, movie model.Netflix) (*mongo.InsertOneResult, error) {
	result, err := collection.InsertOne(ctx, movie)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("internal server error")
	}

	return result, nil
}

func (r *repository) UpdateOneSeries(ctx context.Context, seriesId string) *mongo.UpdateResult {
	objectId, err := primitive.ObjectIDFromHex(seriesId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (r *repository) DeleteOneSeries(ctx context.Context, seriesId string) *mongo.DeleteResult {
	objectId, err := primitive.ObjectIDFromHex(seriesId)

	fmt.Println(objectId)

	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": objectId}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (r *repository) DeleteAllSeries(ctx context.Context) *mongo.DeleteResult {
	result, err := collection.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (r *repository) GetAllSeries(ctx context.Context) []primitive.M {
	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close(ctx)

	var allSeries []primitive.M

	// if err = result.All(ctx, &allSeries); err != nil {
	// 	log.Fatal(err)
	// }

	for result.Next(ctx) {
		var series bson.M
		if err = result.Decode(&series); err != nil {
			log.Fatal(err)
		}
		allSeries = append(allSeries, series)
	}

	return allSeries
}
