package controllers

import (
	"net/http"

	"github.com/darienkentanu/test-erajaya/lib/database"
	"github.com/darienkentanu/test-erajaya/models"

	"github.com/labstack/echo/v4"
)

type M map[string]interface{}

type ProductController struct {
	db database.ProductModel
}

func NewProductController(db database.ProductModel) *ProductController {
	return &ProductController{db: db}
}

func (pc *ProductController) AddProduct(c echo.Context) error {
	var product models.Product

	if err := c.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	product, err := pc.db.AddProduct(product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusCreated, M{
		"status": "success",
		"data":   product,
	})
}

func (pc *ProductController) GetNewestProduct(c echo.Context) error {
	var products []models.Product

	products, err := pc.db.GetNewestProduct()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, M{
		"status": "success",
		"data":   products,
	})
}

func (pc *ProductController) GetProductASC(c echo.Context) error {
	var products []models.Product

	products, err := pc.db.GetProductASC()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, M{
		"status": "success",
		"data":   products,
	})
}

func (pc *ProductController) GetProductDESC(c echo.Context) error {
	var products []models.Product

	products, err := pc.db.GetProductDESC()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, M{
		"status": "success",
		"data":   products,
	})
}

func (pc *ProductController) GetProductHighestPrice(c echo.Context) error {
	var products []models.Product

	products, err := pc.db.GetProductHighestPrice()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, M{
		"status": "success",
		"data":   products,
	})
}

func (pc *ProductController) GetProductLowestPrice(c echo.Context) error {
	var products []models.Product

	products, err := pc.db.GetProductLowestPrice()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, M{
		"status": "success",
		"data":   products,
	})
}
