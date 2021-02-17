package main

import (
	"fmt"
	"net/http"

	"hq-collections.com/collection"
	"hq-collections.com/cors"
	"hq-collections.com/volume"

	"github.com/gorilla/mux"
)

// SetupRoutes builds the routes for the API
func SetupRoutes(apiBasePath string) {
	collectionsHandler := http.HandlerFunc(collection.HandleCollections)
	collectionHandler := http.HandlerFunc(collection.HandleCollection)
	collectionVolumesHandler := http.HandlerFunc(collection.HandleCollectionVolumes)
	volumesHandler := http.HandlerFunc(volume.HandleVolumes)
	volumeHandler := http.HandlerFunc(volume.HandleVolume)

	router := mux.NewRouter()
	router.Use(cors.Middleware)
	apiRouter := router.PathPrefix("/api").
		Methods("GET", "POST", "PUT", "DELETE", "OPTIONS").
		Subrouter()

	apiRouter.HandleFunc("/collections/{id}", collectionHandler)
	apiRouter.HandleFunc("/collections/{id}/volumes", collectionVolumesHandler)
	apiRouter.HandleFunc("/collections", collectionsHandler)

	apiRouter.HandleFunc("/volumes/{id}", volumeHandler)
	apiRouter.HandleFunc("/volumes", volumesHandler)

	http.Handle("/", router)

	fmt.Print("Available routes:\n\n")
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, _ := route.GetPathTemplate()
		met, _ := route.GetMethods()
		fmt.Println(tpl, "", met, "")
		return nil
	})
}
