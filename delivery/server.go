package delivery

import (
	"gin_produk/config"
	"gin_produk/delivery/controller"
	"gin_produk/repo"
	"gin_produk/usecase"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	productUc  usecase.CreateProductUsecase
	listProduk usecase.AllProductUsecase
	engine     *gin.Engine
	host       string
}

//nambain run biar bisa diakses dari luar

func (a *appServer) iniController() {
	controller.NewProductController(a.engine, a.productUc, a.listProduk)
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

	listProdukUc := usecase.NewAllProductUsecase(productRepo)

	// apiHost := os.Getenv("API_HOST")
	//host localhost
	// apiPort := os.Getenv("API_PORT")
	//default port 8888

	c := config.NewConfig()
	host := c.Url

	// host := fmt.Sprintf("%s:%s", apiHost, apiPort)

	return &appServer{
		productUc:  productUc,
		listProduk: listProdukUc,
		engine:     r,
		host:       host,
	}
}
