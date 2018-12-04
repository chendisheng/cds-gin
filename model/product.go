package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type Product struct {
	Id int `json:"id"`
	Code string `json:"code"`
	Price int  `json:"price"`
}

var ProductDao =Product{}

func(p *Product) Insert(product Product) error{

	db, err := gorm.Open("mysql", "root:cds0405@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()
	db.Create(&product)
	return nil

}

