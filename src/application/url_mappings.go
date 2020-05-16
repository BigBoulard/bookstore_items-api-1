package application

import (
	"net/http"

	"github.com/federicoleon/bookstore_items-api/src/controllers"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)

	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controllers.ItemsController.Get).Methods(http.MethodGet)
	router.HandleFunc("/items/{id}", controllers.ItemsController.Upsert).Methods(http.MethodPost)
	router.HandleFunc("/items/search", controllers.ItemsController.Search).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controllers.ItemsController.Delete).Methods(http.MethodDelete)
}
