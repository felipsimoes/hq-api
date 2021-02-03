package volume

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"hq-collections.com/cors"
)

const volumesPath = "volumes"

func handleVolumes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		volumeList := getVolumesList()
		j, err := json.Marshal(volumeList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		var volume Volume
		err := json.NewDecoder(r.Body).Decode(&volume)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = addOrUpdateVolume(volume)
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

func handleVolume(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", volumesPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	volumeID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		volume := getVolume(volumeID)
		if volume == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(volume)
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
		var volume Volume
		err := json.NewDecoder(r.Body).Decode(&volume)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if volume.ID != volumeID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = addOrUpdateVolume(volume)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case http.MethodDelete:
		removeVolume(volumeID)

	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// SetupRoutes is a temporary routing method. It must be improved
// with Mux.
func SetupRoutes(apiBasePath string) {
	volumesHandler := http.HandlerFunc(handleVolumes)
	volumeHandler := http.HandlerFunc(handleVolume)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, volumesPath), cors.Middleware(volumesHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, volumesPath), cors.Middleware(volumeHandler))
}
