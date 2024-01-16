package routes

import (
	"my_projects/pkg/controllers"

	"github.com/gorilla/mux"
)

var CrudRoutes = func(router *mux.Router) {
	router.HandleFunc("/tailor/", controllers.CreateTailor).Methods("POST")
	router.HandleFunc("/tailor/", controllers.GetTailor).Methods("GET")
	router.HandleFunc("/tailor/{id}", controllers.GetTailorById).Methods("GET")
	router.HandleFunc("/tailor/{id}", controllers.UpdateTailor).Methods("PUT")
	router.HandleFunc("/tailor/{id}", controllers.DeleteTailor).Methods("DELETE")
}
