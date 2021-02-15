package volume

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
)

// VolumesMap stores the data for the volumes.
// volumeID is the key for the map, allowing us to access
// volumes without having to iterate over all the items in the
// slice. The value is the volume itself.
//
// We use mutex because our webservice is multi-threaded and maps
// in Golang are not naturally thread-safe, which means we need to wrap
// our map in a mutex to avoid two threads from writing and reading
// the `map` at the same time.
var VolumesMap = struct {
	sync.RWMutex
	m map[int]Volume
}{m: make(map[int]Volume)}

func init() {
	fmt.Print("loading volumes...")
	volumeMap, err := loadVolumesMap()
	VolumesMap.m = volumeMap
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d volumes loaded\n\n", len(VolumesMap.m))
}

func Volumes() []Volume {
	return getVolumesList()
}

func loadVolumesMap() (map[int]Volume, error) {
	fileName := "data/volumes.json"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file [%s] does not exist", fileName)
	}

	file, _ := ioutil.ReadFile(fileName)
	volumeList := make([]Volume, 0)
	err = json.Unmarshal([]byte(file), &volumeList)

	if err != nil {
		log.Fatal(err)
	}

	VolumesMap := make(map[int]Volume)
	for i := 0; i < len(volumeList); i++ {
		VolumesMap[volumeList[i].ID] = volumeList[i]
	}

	return VolumesMap, nil
}

func getVolume(volumeID int) *Volume {
	VolumesMap.RLock()
	defer VolumesMap.RUnlock()

	if volume, ok := VolumesMap.m[volumeID]; ok {
		return &volume
	}

	return nil
}

func removeVolume(volumeID int) {
	VolumesMap.Lock()
	defer VolumesMap.Unlock()
	delete(VolumesMap.m, volumeID)
}

func getVolumesList() []Volume {
	VolumesMap.RLock()
	volumes := make([]Volume, 0, len(VolumesMap.m))
	for _, value := range VolumesMap.m {
		volumes = append(volumes, value)
	}
	VolumesMap.RUnlock()
	return volumes
}

func getVolumesIDs() []int {
	VolumesMap.RLock()
	volumesIDs := []int{}
	for key := range VolumesMap.m {
		volumesIDs = append(volumesIDs, key)
	}
	VolumesMap.RUnlock()
	sort.Ints(volumesIDs)
	return volumesIDs
}

func getNextVolumeID() int {
	volumesIDs := getVolumesIDs()
	return volumesIDs[len(volumesIDs)-1] + 1
}

func addOrUpdateVolume(volume Volume) (int, error) {
	// if the volume id is set, update, otherwise add
	addOrUpdateID := -1
	if volume.ID > 0 {
		oldVolume := getVolume(volume.ID)
		// if it exists, replace it, otherwise return error
		if oldVolume == nil {
			return 0, fmt.Errorf("volume id [%d] doesn't exist", volume.ID)
		}
		addOrUpdateID = volume.ID
	} else {
		addOrUpdateID = getNextVolumeID()
		volume.ID = addOrUpdateID
	}
	VolumesMap.Lock()
	VolumesMap.m[addOrUpdateID] = volume
	VolumesMap.Unlock()
	return addOrUpdateID, nil
}
