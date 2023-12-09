package models

import (
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model

	Title       string  `gorm:"not null" json:"title" binding:"required"`
	Artist      string  `gorm:"not null" json:"artist" binding:"required"`
	ReleaseYear uint    `gorm:"not null" json:"release_year" binding:"required"`
	Genre       string  `gorm:"not null" json:"genre" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `gorm:"not null" json:"price" binding:"required"`
	Status      string  `gorm:"not null" json:"status" binding:"required"`
	SellerName  string  `gorm:"not null" json:"seller_name"`
	ImageUrl    string  `json:"image_url"`
}

type Filter struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

func CreateRecord(db *gorm.DB, title, artist string, genre string, releaseYear uint, description string, price float64, status string, sellerID uint) (*Record, error) {
	seller, err := GetUserByID(db, sellerID)
	if err != nil {
		return nil, err
	}
	record := &Record{
		Title:       title,
		Artist:      artist,
		ReleaseYear: releaseYear,
		Genre:       genre,
		Description: description,
		Price:       price,
		Status:      status,
		SellerName:  seller.Username,
	}

	if err := db.Create(record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func UpdateRecord(db *gorm.DB, recordID uint, title, artist string, genre string, releaseYear uint, description string, price float64, status string) (*Record, error) {
	record, err := GetRecordByID(db, recordID)
	if err != nil {
		return nil, err
	}

	record.Title = title
	record.Artist = artist
	record.ReleaseYear = releaseYear
	record.Genre = genre
	record.Description = description
	record.Price = price
	record.Status = status

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
func DeleteAllRecordsBySellerID(db *gorm.DB, sellerID uint) error {
	var records []Record
	if err := db.Where("seller_id = ?", sellerID).Find(&records).Error; err != nil {
		return err
	}

	if err := db.Delete(&records).Error; err != nil {
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

func GetAllRecords(db *gorm.DB, filters []Filter) ([]Record, error) {
	var records []Record

	query := db

	for _, filter := range filters {
		query = query.Where(filter.Field, filter.Value)
	}

	if err := query.Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func GetRecordsBySellerName(db *gorm.DB, filters []Filter, sellerName string) ([]Record, error) {
	var records []Record
	query := db.Where("seller_name = ?", sellerName)

	for _, filter := range filters {
		query = query.Where(filter.Field, filter.Value)
	}

	if err := query.Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}
func GetSellerName(db *gorm.DB, RecordID uint) (string, error) {
	var record Record
	if err := db.First(&record, RecordID).Error; err != nil {
		return "", err
	}
	return record.SellerName, nil
}

func SearchRecordsWithPagination(db *gorm.DB, filters []Filter, searchTerm string, page, perPage int) ([]Record, error) {
	var records []Record

	offset := (page - 1) * perPage
	query := db
	for _, filter := range filters {
		query = query.Where(filter.Field, filter.Value)
	}
	if err := db.Where("MATCH(title, artist, description) AGAINST (?)", searchTerm).
		Limit(perPage).
		Offset(offset).
		Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}
