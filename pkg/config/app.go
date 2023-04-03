package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //underscore is the requirement from the package side, read their documentation from google if required
)

var (
	db *gorm.DB //variable named db is created to interact with database
)

func Connect() {
	d, err := gorm.Open("mysql", "manu:porwal@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d

}
func GetDB() *gorm.DB {
	return db
}
