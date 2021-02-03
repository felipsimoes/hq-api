package main

import (
	"fmt"
	"net/http"

	"hq-collections.com/collection"
	"hq-collections.com/cors"
	"hq-collections.com/volume"
)

const collectionsPath = "collections"
const volumesPath = "volumes"

// SetupRoutes is a temporary routing method. It must be improved
// with Mux.
func SetupRoutes(apiBasePath string) {
	collectionsHandler := http.HandlerFunc(collection.HandleCollections)
	collectionHandler := http.HandlerFunc(collection.HandleCollection)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, collectionsPath), cors.Middleware(collectionsHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, collectionsPath), cors.Middleware(collectionHandler))

	volumesHandler := http.HandlerFunc(volume.HandleVolumes)
	volumeHandler := http.HandlerFunc(volume.HandleVolume)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, volumesPath), cors.Middleware(volumesHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, volumesPath), cors.Middleware(volumeHandler))

}
