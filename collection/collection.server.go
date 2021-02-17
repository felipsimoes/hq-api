package collection

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandleCollections(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Println("/GET collections")
		collections := FindAllCollections()

		j, err := json.Marshal(collections)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		fmt.Println("/POST collections")
		var collection Collection
		err := json.NewDecoder(r.Body).Decode(&collection)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if collection.Save() {
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
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
		fmt.Println("/GET collection ", collectionID)
		collection, err := GetCollection(collectionID)
		if err != nil {
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
		fmt.Println("/PUT collection ", collectionID)
		var collection Collection
		err := json.NewDecoder(r.Body).Decode(&collection)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if collection.Save() {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	case http.MethodDelete:
		fmt.Println("/DELETE collection ", collectionID)
		collection, err := GetCollection(collectionID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		collection.Destroy()
		w.WriteHeader(http.StatusNoContent)
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
	// collectionID, err := strconv.Atoi(params["id"])
	// if err != nil {
	// 	log.Print(err)
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// }
	switch r.Method {
	case http.MethodGet:
		// collection := getCollection(collectionID)
		// if collection == nil {
		// 	w.WriteHeader(http.StatusNotFound)
		// 	return
		// }
		// j, err := json.Marshal(collection.getCollectionVolumes())
		// if err != nil {
		// 	log.Print(err)
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	return
		// }
		// _, err = w.Write(j)
		// if err != nil {
		// 	log.Fatal(err)
		// }
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
