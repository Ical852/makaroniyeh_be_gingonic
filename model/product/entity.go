package product

import "makaroniyeh/model/category"

type Product struct {
	ID 			int
	Name 		string
	CategoryID 	int
	Desc 		string
	Price 		int
	Image 		string
	Category	category.Category
}