package main

import (
	"gin_produk/model"
	"gin_produk/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	routerEngine := gin.Default()

	routerEngine.GET("/", func(c *gin.Context) {
		c.String(200, "Healthy Check")
	})

	repoCat := repository.NewCategoryRepo()

	routerEngine.GET("/customerPost", func(c *gin.Context) {

	})

	//category GET Method

	//http://localhost:8080/categoryPost

	routerEngine.POST("/categoryPost", func(c *gin.Context) {

		newCategory := model.Category{
			CategoryId:   "C0001",
			CategoryName: "Nasi Bakar",
		}

		dataCategory, err := repoCat.Add(&newCategory)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
				"data":    &dataCategory,
			})
		}

	})

	//Category GET METHOD

	routerEngine.GET("/categoryGet", func(c *gin.Context) {

		categoryList := repoCat.Retrieve()

		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    &categoryList,
		})

	})

	//Post produk
	// http://localhost:8080/produkPost

	repoProd := repository.NewProductRepo()

	routerEngine.POST("/produkPost", func(c *gin.Context) {

		var newProduk model.Product

		// newProduct := model.Product{
		// 	ProductId:   "001",
		// 	ProductName: "Cemilan",
		// 	Category: &model.Category{
		// 		ProductId:   "002",
		// 		ProductName: "Ikan Bakar",
		// 	},
		// }

		// dataCustomer, err := repoCat.Add(&newProduct)

		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"error": err.Error(),
		// 	})
		// } else {
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"message": "OK",
		// 		"data":    &newProduct,
		// 	})
		// }

		if err := c.ShouldBindJSON(&newProduk); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{
				"id":         newProduk.ProductId,
				"name":       newProduk.ProductName,
				"categories": newProduk.Category,
				"is_active":  newProduk.IsActive,
			})
			//nambain produk

			repoProd.Add(newProduk)

			//display produk all
			displayAllProduk := repoProd.RetrieveData()
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
				"data":    &displayAllProduk,
			})
		}

	})

	routerEngine.Run()

}
