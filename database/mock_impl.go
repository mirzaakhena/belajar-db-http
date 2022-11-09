package database

import (
	"belajar/model"
	"database/sql"
)

type MockImpl struct {
}

func NewMockImpl(db *sql.DB) (*MockImpl, error) {
	return &MockImpl{}, nil
}

func (m MockImpl) InsertProduct(product *model.Product) error {
	return nil
}
func (m MockImpl) SelectAllProduct(page, size int) ([]model.Product, error) {
	return nil, nil
}
func (m MockImpl) SelectOneProduct(id int) (*model.Product, error) {
	return nil, nil
}
func (m MockImpl) UpdateOneProduct(product model.Product) error {
	return nil
}
func (m MockImpl) DeleteProduct(id int) error {
	return nil
}
