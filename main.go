package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Collection is a struct that does something
type Collection struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Publisher string `json:"published"`
}

type Volume struct {
	ID           int    `json:"id"`
	CollectionID int    `json:"collection_id"`
	Name         string `json:"name"`
	Edition      string `json:"edition"`
	Pages        int    `json:"pages"`
}

var collectionList = []Collection{
	{1, "Turma da Mônica", "X"},
	{2, "Asterix", "Y"},
}

var volumesList = []Volume{
	{1, 1, "Almanacão", "1", 39},
	{2, 1, "Especial de natal", "especial", 20},
	{3, 2, "Asterix em São Paulo", "3", 120},
	{4, 2, "Asterix en français", "francesa", 2},
}

func collectionHandler(w http.ResponseWriter, r *http.Request) {
	jsonList, err := json.Marshal(collectionList)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(jsonList))
}

// FindVolumes is a function that...
func (collection Collection) FindVolumes() []Volume {
	result := []Volume{}
	for _, volume := range volumesList {
		if volume.CollectionID == collection.ID {
			result = append(result, volume)
		}
	}
	return result
}

func volumesHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("asdasd")
	// urlPathSegments := strings.Split(r.URL.Path, "/volumes")
	// fmt.Println(string(urlPathSegments[0][1:]))
	// collectionID, err := strconv.Atoi(string(urlPathSegments[0][1:]))
	// if err != nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// }

	collectionID := 2

	emptyCollection := Collection{}
	currentCollection := emptyCollection
	for _, item := range collectionList {
		if item.ID == collectionID {
			currentCollection = item
		}
	}

	if currentCollection == emptyCollection {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonList, err := json.Marshal(currentCollection.FindVolumes())
	if err != nil {
		panic(err)
	}

	w.Write([]byte(jsonList))
}

func main() {
	// http.HandleFunc(`/{id:[0-9]+}/volumes`, volumesHandler)
	http.HandleFunc(`/volumes`, volumesHandler)
	http.HandleFunc("/collections", collectionHandler)
	fmt.Println("Waiting for requests...")
	http.ListenAndServe(":5000", nil)
}
