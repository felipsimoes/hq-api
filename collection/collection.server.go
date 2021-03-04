package collection

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandleCollections(w http.ResponseWriter, r *http.Request) {
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

func HandleCollection(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params["id"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	collectionID, err := strconv.Atoi(params["id"])
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

func HandleCollectionVolumes(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params["id"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	collectionID, err := strconv.Atoi(params["id"])
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
		j, err := json.Marshal(collection.getCollectionVolumes())
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
