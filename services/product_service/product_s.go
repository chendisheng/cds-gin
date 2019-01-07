package product_service

import "cds-gin/models"

type ProductService struct {
	Id            int
	Code         string
	Price        float32
	PageNum  int
	PageSize int
}

func (s *ProductService)Add() (err error){

	product := models.Product{
		Code:s.Code,
		Price:s.Price,
	}

	if err := models.AddProduct(product); err != nil {
		return err
	}
	return
}
