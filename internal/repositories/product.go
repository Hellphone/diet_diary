package repositories

import (
	"log"

	"diet_diary/internal/database"
	"diet_diary/internal/domain"

	sq "github.com/Masterminds/squirrel"
)

const tableProduct = "products"

func GetProductById(id int64) (*domain.Product, error) {
	query := `SELECT * FROM products WHERE id = $1`

	product := database.Product{}
	err := database.DB.Get(&product, query, id)
	if err != nil {
		log.Println("Error getting product by id:", err)
		return nil, err
	}

	return database.ProductToDomain(&product), nil
}

func GetProducts(filter *database.Filter) (domain.Products, error) {
	var err error
	builder := sq.Select("*").From(tableProduct)
	builder, err = SetFilters(builder, filter)
	if err != nil {
		log.Println("Error applying filters:", err)
		return nil, err
	}

	products := database.Products{}
	query, args, err := builder.ToSql()
	if err != nil {
		log.Println("Error building SQL query:", err)
		return nil, err
	}

	err = database.DB.Select(&products, rebind(query), args...)
	if err != nil {
		log.Println("Error getting products:", err)
		return nil, err
	}

	return database.ProductsToDomain(products), nil
}

func InsertProduct(product *domain.Product) (int64, error) {
	query := `INSERT INTO products (name, protein, fat, carbohydrate, calories) 
              VALUES (:name, :protein, :fat, :carbohydrate, :calories) RETURNING id`

	rows, err := database.DB.NamedQuery(query, product)
	if err != nil {
		log.Println("Error inserting product:", err)
		return 0, err
	}
	defer rows.Close()

	var id int64
	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func UpdateProduct(product *domain.Product) (int64, error) {
	query := `UPDATE products
              SET name = :name, protein = :protein, fat = :fat, carbohydrate = :carbohydrate, calories = :calories
			  WHERE id = :id`

	rows, err := database.DB.NamedQuery(query, product)
	if err != nil {
		log.Println("Error updating product:", err)
		return 0, err
	}
	defer rows.Close()

	return product.ID, nil
}

func DeleteProduct(id int64) (int64, error) {
	query := `DELETE FROM products
	 		  WHERE id = $1`

	_, err := database.DB.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting product with id = %d: %v", id, err)
		return 0, err
	}

	return id, nil
}
