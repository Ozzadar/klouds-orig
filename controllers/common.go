package controllers

import (

 	_ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "github.com/superordinate/klouds2.0/models"
    "fmt"
)

type ErrorMessage struct {
	Message	string

}

var (
	db *gorm.DB
)

// connect to the db

func InitDB() {

	fmt.Println("Initializing Database connection.")

    dbm, err := gorm.Open("mysql", "root:diamond11@(127.0.0.1:3306)/klouds?charset=utf8&parseTime=True")

    if(err != nil){
        panic("Unable to connect to the database")
    } else {
    	fmt.Println("Database connection established.")
    }

    db = &dbm
    dbm.DB().Ping()
    dbm.DB().SetMaxIdleConns(10)
    dbm.DB().SetMaxOpenConns(100)
    db.LogMode(true)
 
    if !dbm.HasTable(&models.User{}){
        dbm.CreateTable(&models.User{})
    }
}

func CreateUser(u *models.User) {
	fmt.Println("Creating user: " + u.Username)

	db.Create(&u)
}