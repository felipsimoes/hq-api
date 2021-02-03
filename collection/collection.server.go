package collection

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"hq-collections.com/cors"
)

const collectionsPath = "collections"

func handleCollections(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		collectionList := getCollectionsList()
		j, err := json.Marshal(collectionList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		var collection Collection
		err := json.NewDecoder(r.Body).Decode(&collection)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = addOrUpdateCollection(collection)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleCollection(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", collectionsPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	collectionID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		collection := getCollection(collectionID)
		if collection == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(collection)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}

	case http.MethodPut:
		var collection Collection
		err := json.NewDecoder(r.Body).Decode(&collection)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if collection.ID != collectionID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = addOrUpdateCollection(collection)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case http.MethodDelete:
		removeCollection(collectionID)

	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// SetupRoutes is a temporary routing method. It must be improved
// with Mux.
func SetupRoutes(apiBasePath string) {
	collectionsHandler := http.HandlerFunc(handleCollections)
	collectionHandler := http.HandlerFunc(handleCollection)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, collectionsPath), cors.Middleware(collectionsHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, collectionsPath), cors.Middleware(collectionHandler))
}
