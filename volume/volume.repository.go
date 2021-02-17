package volume

import (
	"errors"

	"gorm.io/gorm"
	"hq-collections.com/database"
)

// Destroy removes the volume from the database
// It always return true because Postgresql does
// not return an error when the delete command
// fails.
func (volume Volume) Destroy() bool {
	database.DB.Delete(&volume)
	return true
}

// Save the volume in the database. If the
// volume has an ID, it will be updated.
// Otherwise, it will be created.
// It returns true if the operation succeeds,
// of false if it does not.
func (volume Volume) Save() bool {
	var err error
	if volume.ID > 0 {
		err = database.DB.Create(&volume).Error
	} else {
		err = database.DB.Save(&volume).Error
	}

	return err == nil
}

// FindAllVolumes returns a list with
// all volumes in the database.
func FindAllVolumes() []Volume {
	var volumeList []Volume
	database.DB.Find(&volumeList)
	return volumeList
}

// GetVolume finds a volume in the database
// given its ID. This method returns and error if
// the volume is not found.
func GetVolume(volumeID int) (*Volume, error) {
	var volume Volume
	err := database.DB.First(&volume, volumeID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	return &volume, nil
}
