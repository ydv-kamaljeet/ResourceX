package models

import "gorm.io/gorm"

// var Shelf []Book

type Book struct {
	gorm.Model
	Name    string `json:"name"`
	Author  string `json:"author"`
	Price   int    `json:"price"`
	FileURL string `json:"fileURL"`
}

// type UpdatedBook struct {
// Name  string `json:"name"`
// Price int    `json:"price"`
// }
