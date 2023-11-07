package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model

	RecordID uint   `gorm:"unique;not null" json:"record-id" binding:"required"`
	Url      string `gorm:"unique;not null" json:"url" binding:"required"`
}

func CreateImage(db *gorm.DB, recordID uint, url string) (*Image, error) {
	image := &Image{
		RecordID: recordID,
		Url:      url,
	}

	if err := db.Create(image).Error; err != nil {
		return nil, err
	}

	return image, nil
}

func UpdateImage(db *gorm.DB, imageID uint, recordID uint, url string) (*Image, error) {
	image, err := GetImageByID(db, imageID)
	if err != nil {
		return nil, err
	}

	image.RecordID = recordID
	image.Url = url

	if err = db.Save(image).Error; err != nil {
		return nil, err
	}

	return image, nil
}

func GetImageByID(db *gorm.DB, imageID uint) (*Image, error) {
	var image *Image
	if err := db.First(image, imageID).Error; err != nil {
		return nil, err
	}

	return image, nil
}

func DeleteImage(db *gorm.DB, imageID uint) error {
	var image Image
	if err := db.First(&image, imageID).Error; err != nil {
		return err
	}
	if err := db.Delete(&image).Error; err != nil {
		return err
	}

	return nil
}
