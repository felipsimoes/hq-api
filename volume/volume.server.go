package volume

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func HandleVolumes(w http.ResponseWriter, r *http.Request) {
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

func HandleVolume(w http.ResponseWriter, r *http.Request) {
	// TODO: the `volumes` here is the volumesPath set on routes.go, it is
	// duplicated. This will be fixed when we start using mux.
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", "volumes"))
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
