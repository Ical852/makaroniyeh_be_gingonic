package product

import (
	"github.com/gin-gonic/gin"
	"makaroniyeh/helper"
	"net/http"
)

type productController struct {
	service Service
}

func NewProductController(service Service) *productController {
	return &productController{service}
}

func (co *productController) GetAllProduct(c *gin.Context) {
	products, err := co.service.GetAll()
	if err != nil {
		errorMessage := gin.H{"error" : err.Error()}
		response := helper.APIResponse("Failed to get Products", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formattedProducts := FormatProducts(products)
	response := helper.APIResponse("Success to get Products", http.StatusOK, "success", formattedProducts)

	c.JSON(http.StatusOK, response)
}

func (co *productController) GetProductByCategory(c *gin.Context)  {
	var input GetByCatInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get Products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	products, err := co.service.GetByCategory(input)
	if err != nil {
		errorMessage := gin.H{"error" : err.Error()}
		response := helper.APIResponse("Failed to get Products", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formattedProducts := FormatProducts(products)
	response := helper.APIResponse("Success to get Products by Category", http.StatusOK, "success", formattedProducts)

	c.JSON(http.StatusOK, response)
}

func (co *productController) GetBestSellerProduct(c *gin.Context) {
	products, err := co.service.GetBestSeller()
	if err != nil {
		errorMessage := gin.H{"error" : err.Error()}
		response := helper.APIResponse("Failed to get Products", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formattedProducts := FormatProducts(products)
	response := helper.APIResponse("Success to get Best Seller Products", http.StatusOK, "success", formattedProducts)

	c.JSON(http.StatusOK, response)
}

func (co *productController) GetSuperSpecialProduct(c *gin.Context) {
	products, err := co.service.GetSuperSpesial()
	if err != nil {
		errorMessage := gin.H{"error" : err.Error()}
		response := helper.APIResponse("Failed to get Products", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formattedProducts := FormatProducts(products)
	response := helper.APIResponse("Success to get Super Special Products", http.StatusOK, "success", formattedProducts)

	c.JSON(http.StatusOK, response)
}

func (co *productController) GetPopularProduct(c *gin.Context) {
	products, err := co.service.GetPopular()
	if err != nil {
		errorMessage := gin.H{"error" : err.Error()}
		response := helper.APIResponse("Failed to get Products", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formattedProducts := FormatProducts(products)
	response := helper.APIResponse("Success to get Super Special Products", http.StatusOK, "success", formattedProducts)

	c.JSON(http.StatusOK, response)
}

func (co *productController) GetProductByName(c *gin.Context)  {
	var input GetByNameInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get Products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	products, err := co.service.GetByName(input)
	if err != nil {
		errorMessage := gin.H{"error" : err.Error()}
		response := helper.APIResponse("Failed to get Products", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formattedProducts := FormatProducts(products)
	response := helper.APIResponse("Success to get Products by Name", http.StatusOK, "success", formattedProducts)

	c.JSON(http.StatusOK, response)
}