package product_service

import "cds-gin/models"

type Product struct {
	Id            int
	Code         string
	Price        float32
	PageNum  int
	PageSize int
}

func (* Product)AddProduct() (err error){

	product := models.Product{

	}

	if err := models.AddProduct(product); err != nil {
		return err
	}
	return
}
