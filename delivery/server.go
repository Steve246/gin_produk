package delivery

import (
	"gin_produk/config"
	"gin_produk/delivery/controller"
	"gin_produk/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	// productUc  usecase.CreateProductUsecase
	// listProduk usecase.AllProductUsecase

	useCaseManager manager.UseCaseManager

	engine *gin.Engine
	host   string
}

//nambain run biar bisa diakses dari luar

func Server() *appServer {
	r := gin.Default()

	repoManager := manager.NewRepositoryManager()

	useCaseManager := manager.NewUseCaseManager(repoManager)

	// apiHost := os.Getenv("API_HOST")
	//host localhost
	// apiPort := os.Getenv("API_PORT")
	//default port 8888

	c := config.NewConfig()
	host := c.Url

	// host := fmt.Sprintf("%s:%s", apiHost, apiPort)

	return &appServer{
		useCaseManager: useCaseManager,
		engine:         r,
		host:           host,
	}
}

func (a *appServer) iniController() {
	controller.NewProductController(a.engine, a.useCaseManager.CreateProductUseCase(), a.useCaseManager.ListProductUseCase())
}

func (a *appServer) Run() {
	a.iniController()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
