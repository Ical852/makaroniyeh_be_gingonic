package main

import (
	"github.com/gin-gonic/gin"
	"makaroniyeh/helper"
	"makaroniyeh/model/category"
	"makaroniyeh/model/product"
)

func main() {
	db, err := helper.DbConnect()
	helper.ErrorHelper(err)

	categoryRepository := category.NewRepository(db)
	categoryService := category.NewService(categoryRepository)
	categoryController := category.NewCategoryController(categoryService)

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productController := product.NewProductController(productService)

	router := gin.Default()
	router.Use(CORSMiddleware())

	api := router.Group("api/")

	router.Static("/images", "./images")

	api.GET("/categories", categoryController.GetCategories)
	api.GET("/products", productController.GetAllProduct)
	api.GET("/products/:id", productController.GetProductByCategory)
	api.GET("/products/best", productController.GetBestSellerProduct)
	api.GET("/products/special", productController.GetSuperSpecialProduct)
	api.GET("/products/popular", productController.GetPopularProduct)
	api.GET("/products/name/:name", productController.GetProductByName)

	router.Run(":8000")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}