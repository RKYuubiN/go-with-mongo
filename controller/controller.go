package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rkyuubin/gowithmongo/model"
	"github.com/rkyuubin/gowithmongo/mongo"
)

type Controller interface {
	GetAllSeries(w http.ResponseWriter, r *http.Request)
	UpdateWatched(w http.ResponseWriter, r *http.Request)
	InsertSeries(w http.ResponseWriter, r *http.Request)
	DeleteOneSeries(w http.ResponseWriter, r *http.Request)
	DeleteAllSeries(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	logger      *log.Logger
	respository mongo.Repository
}

func NewController(logger *log.Logger, repository mongo.Repository) Controller {
	return &controller{
		logger:      logger,
		respository: repository,
	}
}

func (c *controller) GetAllSeries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c.logger.Println("calling method GetAllSereies of repository")
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	response := c.respository.GetAllSeries(ctx)
	defer cancel()
	if response == nil {
		json.NewEncoder(w).Encode("No data found")
	} else {
		json.NewEncoder(w).Encode(response)
	}
}

func (c *controller) UpdateWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	response := c.respository.UpdateOneSeries(ctx, params["id"])
	if response.ModifiedCount > 0 {
		modified := fmt.Sprintf("Updated series. Count %d", response.ModifiedCount)
		json.NewEncoder(w).Encode(modified)
	} else {
		json.NewEncoder(w).Encode("Oops something went wrong")
	}
}

func (c *controller) InsertSeries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var series model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&series)
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	response, err := c.respository.InsertOneSeries(ctx, series)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Internal Server error")
		return
	}
	inserted := fmt.Sprintf("Data inserted of object _id %s", response.InsertedID)
	json.NewEncoder(w).Encode(inserted)
}

func (c *controller) DeleteOneSeries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	response := c.respository.DeleteOneSeries(ctx, params["id"])
	deleted := fmt.Sprintf("Successfully deleted series. Count: %d", response.DeletedCount)
	json.NewEncoder(w).Encode(deleted)
}

func (c *controller) DeleteAllSeries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	response := c.respository.DeleteAllSeries(ctx)
	deleted := fmt.Sprintf("Successfully deleted all series. Count: %d", response.DeletedCount)
	json.NewEncoder(w).Encode(deleted)
}
