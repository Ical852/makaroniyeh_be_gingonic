package product

import "makaroniyeh/model/category"

type ProductFormatter struct {
	ID 			int					`json:"id"`
	Name 		string				`json:"name"`
	CategoryID	int					`json:"category_id"`
	Desc 		string				`json:"desc"`
	Price 		int					`json:"price"`
	Image 		string				`json:"image"`
	Category 	category.CategoryFormatter	`json:"category"`
}

type ProductCategoryFormatter struct {
	ID 		int 	`json:"id"`
	Name 	string 	`json:"name"`
}

func FormatProduct(product Product) ProductFormatter {
	var productFormatted = ProductFormatter{
		ID:         product.ID,
		Name:       product.Name,
		CategoryID: product.CategoryID,
		Desc:       product.Desc,
		Price:      product.Price,
		Image:      product.Image,
		Category:   category.FormatCategory(product.Category),
	}

	return productFormatted
}

func FormatProducts(products []Product) []ProductFormatter {
	var productsFormatted []ProductFormatter
	for _, product := range products{
		productsFormatted = append(productsFormatted, FormatProduct(product))
	}

	return productsFormatted
}