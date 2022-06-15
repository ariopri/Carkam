package repository

import "database/sql"

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *productRepository {
	return &productRepository{db: db}

}

func (p *productRepository) FetchProductByID(id int64) (kampus, error) {
	var kampus Product
	err := p.db.QueryRow("SELECT * FROM products WHERE id = ?", id).Scan(&kampus.ID, &kampus.Name, &kampus.Address, &kampus.Phone, &kampus.Image,
		&kampus.Akreditasi, &kampus.Website)
	if err != nil {
		return kampus, err
	}
	return kampus, nil
}

func (p *ProductRepository) FetchProductByName(productName string) (kampus, error) {
	var kampus Product
	err := p.db.QueryRow("SELECT * FROM products WHERE name = ?", productName).Scan(&kampus.ID, &kampus.Name, &kampus.Address, &kampus.Phone, &kampus.Image,
		&kampus.Akreditasi, &kampus.Website)
	if err != nil {
		return kampus, err
	}
	return kampus, nil
}

func (p *productRepository) FetchProducts() ([]kampus, error) {
	rows, err := p.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kampuss []Product
	for rows.Next() {
		var kampus Product
		err := rows.Scan(&kampus.ID, &kampus.Name, &kampus.Address, &kampus.Phone, &kampus.Image,
			&kampus.Akreditasi, &kampus.Website)
		if err != nil {
			return nil, err
		}
		kampuss = append(kampuss, kampus)
	}
	return kampuss, nil

}
