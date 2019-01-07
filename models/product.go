package models

import "github.com/jinzhu/gorm"

type Product struct {
	Id int `json:"id"`
	Code string `json:"code"`
	Price float32  `json:"price"`
}

func ExistProductByCode(code string) (bool, error) {
	var product Product
	err := db.Select("code").Where("id = ? ", code).First(&product).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if product.Id > 0 {
		return true, nil
	}

	return false, nil
}

func AddProduct(data Product) error {
	product := Product{
		Code:          data.Code,
		Price:         data.Price,
	}
	if err := db.Create(&product).Error; err != nil {
		return err
	}

	return nil
}

func GetProducts(pageNum int, pageSize int, maps interface{}) ([]*Product, error) {
	var products []*Product
	err := db.Preload("code").Where(maps).Offset(pageNum).Limit(pageSize).Find(&products).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return products, nil
}