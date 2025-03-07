package services

import (
	"diet_diary/internal/database"
	"diet_diary/internal/domain"
	"diet_diary/internal/repositories"
)

func GetProducts(filter *database.Filter) (domain.Products, error) {
	return repositories.GetProducts(filter)
}

func GetProductById(id int64) (*domain.Product, error) {
	return repositories.GetProductById(id)
}

func InsertProduct(product *domain.Product) (int64, error) {
	if product.Calories == 0 {
		product.CalcCalories()
	}

	return repositories.InsertProduct(product)
}

func UpdateProduct(product *domain.Product) (int64, error) {
	_, err := GetProductById(product.ID)
	if err != nil {
		return 0, err
	}

	if product.Calories == 0 {
		product.CalcCalories()
	}

	return repositories.UpdateProduct(product)
}

func DeleteProduct(id int64) (int64, error) {
	_, err := GetProductById(id)
	if err != nil {
		return 0, err
	}

	return repositories.DeleteProduct(id)
}
