# Simple CRUD 

This Application using gin as web framework and built-in database/sql to MySQL

We also implement simple layering for every
* model
* database
* server

## Model

Model has simple *Product* which has 4 field

```
type Product struct {
	ID    int64   `json:"id"`
	Nama  string  `json:"nama"`
	Harga float64 `json:"harga"`
	Stok  int     `json:"stok"`
}
```

It also has a Validate method to check whether the field is in correct value

For now we only check for
* nama must not empty
* Stok must more than zero

## Database

For database we simply use 
* Insert Product to DB
* Get All Product from DB
* Get One Specific Product By ID from DB
* Update Product by ID
* Delete (hard delete) Product by ID

## Server

For server we publish all we crud into network
