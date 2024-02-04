package service

import (
	"github.com/stretchr/testify/assert"
	"os"
	"product-app/domain"
	"product-app/service"
	"testing"
)

var productService service.IProductService

func TextMaim(m *testing.M) {
	initialProducts := []domain.Product{
		{
			Id:    1,
			Name:  "AirFryer",
			Price: 1000.0,
			Store: "ABC TECH",
		},
		{
			Id:    2,
			Name:  "UTU",
			Price: 2000.0,
			Store: "ABC TECH",
		},
	}
	FakeProductRepository := NewFakeProductRepository(initialProducts)
	productService = service.NewProductService(FakeProductRepository)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func Test_GetAllProducts(t *testing.T) {
	t.Run("ShouldGetAllProducts", func(t *testing.T) {
		actualProducts := productService.GetAllProducts()
		assert.Equal(t, 2, len(actualProducts))
	})
}
