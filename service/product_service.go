package service

import (
	"errors"
	"product-app/domain"
	"product-app/persistence"
	"product-app/service/model"
)

type IProductService interface {
	Add(productCreate model.ProductCreate) error
	DeleteById(productID int64) error
	GetById(productId int64) (domain.Product, error)
	UpdatePrice(productId int64, newPrice float32) error
	GetAllProducts() []domain.Product
	GetAllProductsByStore(storeName string) []domain.Product
}

type ProductServices struct {
	productRepository persistence.IProductRepository
}

func NewProductService(productRepository persistence.IProductRepository) IProductService {
	return &ProductServices{
		productRepository: productRepository,
	}
}

func (productServices *ProductServices) Add(productCreate model.ProductCreate) error {
	validateErr := validateProductCreate(productCreate)
	if validateErr != nil {
		return validateErr
	}
	return productServices.productRepository.AddProducts(domain.Product{
		Name:     productCreate.Name,
		Price:    productCreate.Price,
		Discount: productCreate.Discount,
		Store:    productCreate.Store,
	})
}

func (productServices *ProductServices) DeleteById(productId int64) error {
	return productServices.productRepository.DeleteById(productId)
}
func (productServices *ProductServices) GetById(productId int64) (domain.Product, error) {
	return productServices.productRepository.GetById(productId)
}
func (productServices *ProductServices) UpdatePrice(productId int64, newPrice float32) error {
	return productServices.productRepository.UpdatePrice(productId, newPrice)
}
func (productServices *ProductServices) GetAllProducts() []domain.Product {
	return productServices.productRepository.GetAllProducts()
}
func (productServices *ProductServices) GetAllProductsByStore(storeName string) []domain.Product {
	return productServices.productRepository.GetAllProductsByStore(storeName)
}

func validateProductCreate(productCreate model.ProductCreate) error {
	if productCreate.Discount > 70.0 {
		return errors.New("Discount can not be great than 70")
	}
	return nil
}
