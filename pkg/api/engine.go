package api

import (
	"database/sql"
	"gfg/pkg/api/product"
	"gfg/pkg/api/seller"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

// CreateAPIEngine creates engine instance that serves API endpoints,
// consider it as a router for incoming requests.
func CreateAPIEngine(db *sql.DB) (*gin.Engine, error) {
	r := gin.New()

	// todo change http url to real url by env variable
	r.Use(location.Default())

	v1 := r.Group("api/v1")

	productRepository := product.NewRepository(db)
	sellerRepository := seller.NewRepository(db)
	emailProvider := seller.NewEmailProvider()
	productController := product.NewController(productRepository, sellerRepository, emailProvider)
	v1.GET("products", productController.List)
	v1.GET("product", productController.Get)
	v1.POST("product", productController.Post)
	v1.PUT("product", productController.Put)
	v1.DELETE("product", productController.Delete)
	sellerController := seller.NewController(sellerRepository)
	v1.GET("sellers", sellerController.List)

	v2 := r.Group("api/v2")
	v2.GET("products", productController.List)
	v2.GET("product", productController.Get)


	return r, nil
}