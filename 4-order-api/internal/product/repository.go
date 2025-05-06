package product

import "4-order-api/pkg/db"

type ProductRepository struct {
	Database *db.Db
}

func NewProductRepository(database *db.Db) *ProductRepository {
	return &ProductRepository{
		Database: database,
	}
}
func (repos *ProductRepository) Create(product *Product) (*Product, error) {
	result := repos.Database.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}
func (repos *ProductRepository) Update(product *Product) (*Product, error) {
	result := repos.Database.DB.Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}
func (repos *ProductRepository) FindId(id string) (*Product, error) {
	var prod Product
	res := repos.Database.DB.First(&prod, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &prod, nil
}
func (repos *ProductRepository) Delete(id string) error {
	res := repos.Database.DB.Delete(&Product{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (repos *ProductRepository) GetId(id string) error {
	var prod Product
	res := repos.Database.DB.First(&prod, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (repos *ProductRepository) GetAllProd() ([]Product, error) {
	var allProd []Product
	res := repos.Database.DB.Find(&allProd)
	if res.Error != nil {
		return nil, res.Error
	}
	return allProd, nil
}
