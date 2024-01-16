package models

import (
	"github.com/HammadShakeel463/go-crud/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Tailor struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Cnic      string `json:"cnic"`
	PhoneNo   string `json:"phoneno"`
	Email     string `json:"email"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Tailor{})
}

func (b *Tailor) CreateTailor() *Tailor {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetTailor() []Tailor {
	var tailors []Tailor
	db.Find(&tailors)
	return tailors
}

func GetTailorById(id int64) (*Tailor, *gorm.DB) {
	var getTailor Tailor
	db := db.Where("ID=?", id).Find(&getTailor)
	return &getTailor, db
}

func DeleteTailor(id int64) Tailor {
	var tailor Tailor
	db.Where("ID=?", id).Delete(&tailor)
	return tailor
}
