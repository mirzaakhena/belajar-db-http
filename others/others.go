package others

import (
	"belajar/model"
	"encoding/json"
	"fmt"
)

func TentangPointer() {
	var b int = 4
	var a *int = &b
	fmt.Printf("%v %v\n", b, &b)
	fmt.Printf("%v %v %v\n", a, &a, *a)

	*a = 7
	fmt.Printf("%v\n", b)

	// 4 0xBBB
	// 0xBBB 0xAAA 4
}

func ProductJSON() {
	productInStr := `[
		{"id":"a","nama":"Jeruk","harga":15000,"stok":8},
		{"id":"b","nama":"Mangga","harga":32000,"stok":3}
	]`

	var newObjects []*model.Product
	err := json.Unmarshal([]byte(productInStr), &newObjects)
	if err != nil {
		panic(err)
	}

	for _, newObject := range newObjects {

		product, err := model.NewProduct(
			newObject.ID,
			newObject.Nama,
			newObject.Harga,
			newObject.Stok,
		)

		if err != nil {
			panic(err)
		}

		// productInBytes, err := json.Marshal(product)
		productInBytes, err := json.MarshalIndent(product, "", "  ")
		if err != nil {
			panic(err)
		}

		_ = productInBytes

		// fmt.Printf("%v\n", string(productInBytes))

	}
}
