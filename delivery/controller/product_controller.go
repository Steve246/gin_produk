package controller

import (
	"gin_produk/model"
	"gin_produk/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	router    *gin.Engine
	ucProduct usecase.CreateProductUsecase
	//nambain buat list
	ucListProduct usecase.AllProductUsecase
}

//nambain retrieve disini

func (p *ProductController) Retrieve(c *gin.Context) {

	products, err := p.ucListProduct.DisplayAll()
	//blm tambain displayAll()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": "Error when creating Product",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"message": products,
	})

}

func (p *ProductController) createNewProduct(c *gin.Context) {
	var newProduct *model.Product
	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := p.ucProduct.CreateProduct(newProduct)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": "Error when creating Product",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": newProduct,
		})

	}

}

func NewProductController(router *gin.Engine, ucProduct usecase.CreateProductUsecase, ucList usecase.AllProductUsecase) *ProductController {
	//disini akan terdapat kumpulan semua request method yang dibutuhkan
	controller := ProductController{
		router:        router,
		ucProduct:     ucProduct,
		ucListProduct: ucList,
	}

	//ini buat taro method-method nya
	router.POST("/product", controller.createNewProduct)

	//taro di retrieve
	router.GET("/productList", controller.Retrieve)

	return &controller
}
