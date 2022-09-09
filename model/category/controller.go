package category

import (
	"github.com/gin-gonic/gin"
	"makaroniyeh/helper"
	"net/http"
)

type categoryController struct {
	service Service
}

func NewCategoryController(service Service) *categoryController {
	return &categoryController{service}
}

func (co *categoryController) GetCategories(c *gin.Context)  {
	campaigns, err := co.service.GetAll()
	if err != nil {
		errorMessage := gin.H{"error" : err.Error()}
		response := helper.APIResponse("Failed to get Categories", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formattedCategories := FormatCategories(campaigns)
	response := helper.APIResponse("Success to get Categories", http.StatusOK, "success", formattedCategories)

	c.JSON(http.StatusOK, response)
}