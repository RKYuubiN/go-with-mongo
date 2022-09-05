package router

import (
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/rkyuubin/gowithmongo/controller"
	"github.com/rkyuubin/gowithmongo/mongo"
)

var (
	control controller.Controller
)

func init() {
	logger := log.New(os.Stdout, "logging", 1)
	repository := mongo.NewRepository(logger)
	control = controller.NewController(logger, repository)
}
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", control.GetAllSeries).Methods("GET")
	r.HandleFunc("/series", control.InsertSeries).Methods("POST")
	r.HandleFunc("/series/{id}", control.UpdateWatched).Methods("PUT")
	r.HandleFunc("/series/{id}", control.DeleteOneSeries).Methods("DELETE")
	r.HandleFunc("/series", control.DeleteAllSeries).Methods("DELETE")
	return r
}
