package collection

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
)

// collectionID is the key for the map, allowing us to access
// collections without having to iterate over all the items in the
// slice. The value is the collection itself.
//
// We use mutex because our webserice is multi-threaded and maps
// in Golang are not naturally thread-safe, which means we need to wrap
// our map in a mutex to avoid two threads from writing and reading
// the `map` at the same time.
var collectionsMap = struct {
	sync.RWMutex
	m map[int]Collection
}{m: make(map[int]Collection)}

func init() {
	fmt.Println("loading collections...")
	collectionMap, err := loadCollectionsMap()
	collectionsMap.m = collectionMap
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("%d collections loaded...\n", len(collectionsMap.m))
}

func loadCollectionsMap() (map[int]Collection, error) {
	fileName := "data/collections.json"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file [%s] does not exist", fileName)
	}

	file, _ := ioutil.ReadFile(fileName)
	collectionList := make([]Collection, 0)
	err = json.Unmarshal([]byte(file), &collectionList)

	if err != nil {
		log.Fatal(err)
	}

	collectionMap := make(map[int]Collection)
	for i := 0; i < len(collectionList); i++ {
		collectionMap[collectionList[i].ID] = collectionList[i]
	}

	return collectionMap, nil
}

func getCollection(collectionID int) *Collection {
	collectionsMap.RLock()
	defer collectionsMap.RUnlock()

	if collection, ok := collectionsMap.m[collectionID]; ok {
		return &collection
	}

	return nil
}

func removeCollection(collectionID int) {
	collectionsMap.Lock()
	defer collectionsMap.Unlock()
	delete(collectionsMap.m, collectionID)
}

func getCollectionsList() []Collection {
	collectionsMap.RLock()
	collections := make([]Collection, 0, len(collectionsMap.m))
	for _, value := range collectionsMap.m {
		collections = append(collections, value)
	}
	collectionsMap.RUnlock()
	return collections
}

func getCollectionsIDs() []int {
	collectionsMap.RLock()
	collectionsIDs := []int{}
	for key := range collectionsMap.m {
		collectionsIDs = append(collectionsIDs, key)
	}
	collectionsMap.RUnlock()
	sort.Ints(collectionsIDs)
	return collectionsIDs
}

func getNextCollectionID() int {
	collectionIDs := getCollectionsIDs()
	return collectionIDs[len(collectionIDs)-1] + 1
}

func addOrUpdateCollection(collection Collection) (int, error) {
	// if the collection id is set, update, otherwise add
	addOrUpdateID := -1
	if collection.ID > 0 {
		oldCollection := getCollection(collection.ID)
		// if it exists, replace it, otherwise return error
		if oldCollection == nil {
			return 0, fmt.Errorf("collection id [%d] doesn't exist", collection.ID)
		}
		addOrUpdateID = collection.ID
	} else {
		addOrUpdateID = getNextCollectionID()
		collection.ID = addOrUpdateID
	}
	collectionsMap.Lock()
	collectionsMap.m[addOrUpdateID] = collection
	collectionsMap.Unlock()
	return addOrUpdateID, nil
}
