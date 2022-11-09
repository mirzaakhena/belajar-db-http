package model

import "fmt"

type Product struct {
	ID    int     `json:"id"`
	Nama  string  `json:"nama"`
	Harga float64 `json:"harga"`
	Stok  int     `json:"stok"`
}

func (p Product) Validate() error {

	if p.Nama == "" {
		return fmt.Errorf("nama tidak boleh kosong")
	}

	if p.Harga <= 0 {
		return fmt.Errorf("stok harus > 0")
	}

	return nil
}

func NewProduct(id int, nama string, harga float64, stok int) (*Product, error) {

	product := Product{
		ID:    id,
		Nama:  nama,
		Harga: harga,
		Stok:  stok,
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return &product, nil
}
