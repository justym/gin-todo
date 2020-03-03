package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	gorm.Model
	Title  string
	Status string
}

type DB struct{}

//Init initialize DB
func (d *DB) Init() error {
	db, err := gorm.Open("sqlite3", "model/DB/test.sqlite3")
	if err != nil {
		return err
	}

	db.AutoMigrate(&Todo{})
	db.Close()

	return err
}

//Insert insert todo into DB
func (d *DB) Insert(title string, status string) error {
	db, err := gorm.Open("sqlite3", "model/DB/test.sqlite3")
	if err != nil {
		return err
	}

	db.Create(&Todo{Title: title, Status: status})
	db.Close()

	return err
}

//Update update DB
func (d *DB) Update(id int, title string, status string) error {
	db, err := gorm.Open("sqlite3", "model/DB/test.sqlite3")
	if err != nil {
		return err
	}

	var item Todo
	db.First(&item, id)
	item.Title = title
	item.Status = status

	db.Save(&item)
	db.Close()

	return err
}

//Delete delete item from DB
func (d *DB) Delete(id int) error {
	db, err := gorm.Open("sqlite3", "model/DB/test.sqlite3")
	if err != nil {
		return err
	}

	var item Todo
	db.First(&item, id)
	db.Delete(&item)
	db.Close()

	return err
}

//All gets All list from DB
func (d *DB) All() ([]Todo, error) {
	db, err := gorm.Open("sqlite3", "model/DB/test.sqlite3")
	if err != nil {
		return nil, err
	}

	var items []Todo
	db.Order("created_at DESC").Find(&items)
	db.Close()

	return items, err
}

//Get get item using id from DB
func (d *DB) Get(id int) (Todo, error) {
	var target Todo
	db, err := gorm.Open("sqlite3", "model/DB/test.sqlite3")
	if err != nil {
		return target, err
	}

	db.First(&target, id)
	db.Close()

	return target, nil
}
