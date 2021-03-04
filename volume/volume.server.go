package volume

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandleVolumes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Println("/GET volumes")
		volumes := FindAllVolumes()

		j, err := json.Marshal(volumes)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		fmt.Println("/POST volumes")
		var volume Volume
		err := json.NewDecoder(r.Body).Decode(&volume)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if volume.Save() {
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

func HandleVolume(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params["id"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	volumeID, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		fmt.Println("/GET volume ", volumeID)
		volume, err := GetVolume(volumeID)
		if err != nil {
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
		fmt.Println("/PUT volume ", volumeID)
		var volume Volume
		err := json.NewDecoder(r.Body).Decode(&volume)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if volume.Save() {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	case http.MethodDelete:
		fmt.Println("/DELETE volume ", volumeID)
		volume, err := GetVolume(volumeID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		volume.Destroy()
		w.WriteHeader(http.StatusNoContent)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
