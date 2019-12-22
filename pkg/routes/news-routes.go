package routes

import (
	"news-2/pkg/controllers"

	"github.com/gorilla/mux"
)

// RegisterNewsStoreRoutes - Endpoint that we can access to run the corresponding functions
var RegisterNewsStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/news/", controllers.CreateNews).Methods("POST")
	router.HandleFunc("/news/", controllers.GetNews).Methods("GET")
	router.HandleFunc("/news/{newsId}", controllers.GetNewsByID).Methods("GET")
	router.HandleFunc("/news/{newsId}", controllers.UpdateNews).Methods("PUT")
	router.HandleFunc("/news/{newsId}", controllers.DeleteNews).Methods("DELETE")
}
