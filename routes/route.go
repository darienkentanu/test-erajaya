package routes

import (
	"github.com/darienkentanu/test-erajaya/controllers"
	"github.com/darienkentanu/test-erajaya/lib/database"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *echo.Echo {
	e := echo.New()
	pdb := database.NewProductDB(db)
	pc := controllers.NewProductController(pdb)

	//tambah product
	e.POST("/products", pc.AddProduct)

	// product terbaru
	e.GET("/newestproducts", pc.GetNewestProduct)
	// product termurah
	e.GET("/productslowest", pc.GetProductLowestPrice)
	// product termahal
	e.GET("/productshighest", pc.GetProductHighestPrice)
	// product a-z
	e.GET("/productsasc", pc.GetProductASC)
	// product z-a
	e.GET("/productsdesc", pc.GetProductDESC)

	return e
}
