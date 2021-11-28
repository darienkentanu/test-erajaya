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

	// product terbaru = newest
	// product termurah = lowest
	// product termahal = highest
	// product a-z = asc
	// product z-a = desc
	e.GET("/products", pc.GetProduct)

	return e
}
