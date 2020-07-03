package repository

import (
	"log"
	"mini-ecommerce/model"
	"os"

	"github.com/jinzhu/gorm"
	//dialect for mysql databae
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DB -> connection to Database
func DB() *gorm.DB {

	db, err := gorm.Open("mysql", os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Error connecting to Database")
		return nil
	}

	db.AutoMigrate(&model.User{}, &model.Product{})
	return db

}
