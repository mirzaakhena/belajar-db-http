package model

import "fmt"

type Product struct {
	ID    int     `json:"id"`
	Nama  string  `json:"nama"`
	Harga float64 `json:"harga"`
	Stok  int     `json:"stok"`
}

func NewProduct(id int, nama string, harga float64, stok int) (*Product, error) {

	if nama == "" {
		return nil, fmt.Errorf("nama tidak boleh kosong")
	}

	if stok <= 0 {
		return nil, fmt.Errorf("stok harus > 0")
	}

	product := Product{
		ID:    id,
		Nama:  nama,
		Harga: harga,
		Stok:  stok,
	}

	return &product, nil
}
