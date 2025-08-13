package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var Db *gorm.DB  
type Question_Schema struct {
	gorm.Model
	QuestionID  uint   `gorm:"primaryKey"`
	Title       string
	Description string
	Question    string
	Answer      []string
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&Question{})
}

func connect_db() {

	dsn := "root:admin@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Print(Db)
	


}

