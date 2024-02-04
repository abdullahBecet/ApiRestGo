package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"product-app/controller/request"
	"product-app/service"
	"strconv"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}
func (productController *ProductController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/products/:id", productController.GetProductById)
	e.GET("/api/v1/products", productController.GetAllProducts)
	e.POST("/api/v1/products", productController.AddProduct)
	e.PUT("/api/v1/products/:id", productController.UpdatePrice)
	e.DELETE("/api/v1/products/:id", productController.DeleteProductById)
}

func (productController *ProductController) GetProductById(c echo.Context) error {
	param := c.Param("id")
	productId, _ := strconv.Atoi(param)

	product, err := productController.productService.GetById(int64(productId))

	if err != nil {
		return c.JSON(http.StatusNotFound, "Not Found")
	}
	return c.JSON(http.StatusOK, product)
}
func (productController *ProductController) GetAllProducts(c echo.Context) error {
	store := c.QueryParam("store")
	if len(store) == 0 {
		allProducts := productController.productService.GetAllProducts()
		return c.JSON(http.StatusOK, allProducts)
	}
	productswithGivenStore := productController.productService.GetAllProductsByStore(store)
	return c.JSON(http.StatusOK, productswithGivenStore)
}
func (productController *ProductController) AddProduct(c echo.Context) error {
	var addProductRequest request.AddProductRequest
	binderr := c.Bind(&addProductRequest)
	if binderr != nil {
		return c.JSON(http.StatusBadRequest, "Has a Problem")
	}
	herErr := productController.productService.Add(addProductRequest.ToModel())
	if herErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, binderr.Error())

	}
	return c.NoContent(http.StatusCreated)
}
func (productController *ProductController) UpdatePrice(c echo.Context) error {
	param := c.Param("id")
	productId, _ := strconv.Atoi(param)
	newPrice := c.QueryParam("NewPrice")
	if len(newPrice) == 0 {
		return c.JSON(http.StatusBadRequest, "Give a new price")
	}
	convertedPrice, _ := strconv.ParseFloat(newPrice, 32)
	productController.productService.UpdatePrice(int64(productId), float32(convertedPrice))
	return c.NoContent(http.StatusOK)
}
func (productController *ProductController) DeleteProductById(c echo.Context) error {
	param := c.Param("id")
	productId, _ := strconv.Atoi(param)

	err := productController.productService.DeleteById(int64(productId))
	if err != nil {
		return c.JSON(http.StatusNotFound, "Not Delete")
	}
	return c.NoContent(http.StatusOK)
}
