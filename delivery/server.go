package delivery

import (
	"fmt"
	"gin_produk/delivery/controller"
	"gin_produk/repo"
	"gin_produk/usecase"
	"os"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	productUc usecase.CreateProductUsecase
	engine    *gin.Engine
	host      string
}

//nambain run biar bisa diakses dari luar

func (a *appServer) iniController() {
	controller.NewProductController(a.engine, a.productUc)
}

func (a *appServer) Run() {
	a.iniController()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}

func Server() *appServer {
	r := gin.Default()

	productRepo := repo.NewProductRepository()

	productUc := usecase.NewCreateProductUseCase(productRepo)

	apiHost := os.Getenv("API_HOST")
	//host localhost
	apiPort := os.Getenv("API_PORT")
	//default port 8888

	host := fmt.Sprintf("%s:%s", apiHost, apiPort)

	return &appServer{
		productUc: productUc,
		engine:    r,
		host:      host,
	}
}
