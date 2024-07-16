package models

import (
	"gorm.io/gorm"
)
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
type DB interface {
    Find(out interface{}, where ...interface{}) *gorm.DB
    Create(value interface{}) *gorm.DB
    First(out interface{}, where ...interface{}) *gorm.DB
    Model(value interface{}) *gorm.DB
    Where(query interface{}, args ...interface{}) *gorm.DB
    Delete(value interface{}, where ...interface{}) *gorm.DB
    AutoMigrate(dst ...interface{}) error
}