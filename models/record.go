package models

import (
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model

	Title       string `gorm:"not null"`
	Artist      string `gorm:"not null"`
	ReleaseYear uint   `gorm:"not null"`
	Description string
	Price       float64 `gorm:"not null"`
	Status      string  `gorm:"not null"`
	SellerID    uint    `gorm:"not null"`
}

func CreateRecord(db *gorm.DB, title, artist string, releaseYear uint, description string, price float64, status string, sellerID uint) (*Record, error) {
	record := &Record{
		Title:       title,
		Artist:      artist,
		ReleaseYear: releaseYear,
		Description: description,
		Price:       price,
		Status:      status,
		SellerID:    sellerID,
	}

	if err := db.Create(record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func UpdateRecord(db *gorm.DB, recordID uint, title, artist string, releaseYear uint, description string, price float64, status string, sellerID uint) (*Record, error) {
	record, err := GetRecordByID(db, recordID)
	if err != nil {
		return nil, err
	}

	record.Title = title
	record.Artist = artist
	record.ReleaseYear = releaseYear
	record.Description = description
	record.Price = price
	record.Status = status
	record.SellerID = sellerID

	if err = db.Save(record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func DeleteRecord(db *gorm.DB, recordID uint) error {
	var record Record
	if err := db.First(&record, recordID).Error; err != nil {
		return err
	}

	if err := db.Delete(&record).Error; err != nil {
		return err
	}

	return nil
}

func GetRecordByID(db *gorm.DB, recordID uint) (*Record, error) {
	var record Record
	if err := db.First(&record, recordID).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

func GetAllRecords(db *gorm.DB) ([]Record, error) {
	var records []Record
	if err := db.Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func GetRecordsBySellerID(db *gorm.DB, sellerID uint) ([]Record, error) {
	var records []Record
	if err := db.Where("seller_id = ?", sellerID).Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func SearchRecordsWithPagination(db *gorm.DB, query string, page, perPage int) ([]Record, error) {
	var records []Record

	offset := (page - 1) * perPage

	if err := db.Where("MATCH(title, artist, description) AGAINST (?)", query).
		Limit(perPage).
		Offset(offset).
		Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}
