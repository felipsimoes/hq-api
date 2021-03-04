package collection

import (
	"errors"

	"gorm.io/gorm"
	"hq-collections.com/database"
)

// Destroy removes the collection from the database
// It always return true because Postgresql does
// not return an error when the delete command
// fails.
func (collection Collection) Destroy() bool {
	database.DB.Delete(&collection)
	return true
}

// Save the collection in the database. If the
// collection has an ID, it will be updated.
// Otherwise, it will be created.
// It returns true if the operation succeeds,
// of false if it does not.
func (collection Collection) Save() bool {
	var err error
	if collection.ID > 0 {
		err = database.DB.Create(&collection).Error
	} else {
		err = database.DB.Save(&collection).Error
	}

	return err == nil
}

// FindAllCollections returns a list with
// all collections in the database.
func FindAllCollections() []Collection {
	var collectionList []Collection
	database.DB.Find(&collectionList)
	return collectionList
}

// GetCollection finds a collection in the database
// given its ID. This method returns and error if
// the collection is not found.
func GetCollection(collectionID int) (*Collection, error) {
	var collection Collection
	err := database.DB.First(&collection, collectionID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	return &collection, nil
}
