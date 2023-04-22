package product

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	dbCommon "github.com/aziemp66/byte-bargain/common/db"
	httpCommon "github.com/aziemp66/byte-bargain/common/http"

	productRepository "github.com/aziemp66/byte-bargain/internal/repository/product"
)

type ProductUsecaseImplementation struct {
	ProductRepository productRepository.Repository
	DB                *sql.DB
}

func NewProductUsecaseImplementation(productRepository productRepository.Repository, db *sql.DB) *ProductUsecaseImplementation {
	return &ProductUsecaseImplementation{
		ProductRepository: productRepository,
		DB:                db,
	}
}

func (p ProductUsecaseImplementation) GetRecommendedProduct(ctx *gin.Context) ([]httpCommon.Product, error) {
	tx, err := p.DB.Begin()

	if err != nil {
		return nil, err
	}

	defer dbCommon.CommitOrRollback(tx)

	orderProducts, err := p.ProductRepository.GetAllOrderProduct(ctx, tx)

	if err != nil {
		return nil, err
	}

	type productCount struct {
		ProductID string
		Count     int
	}

	var productCounts []productCount

	for _, orderProduct := range orderProducts {
		var isExist bool

		for i, productCount := range productCounts {
			if productCount.ProductID == orderProduct.ProductID {
				productCounts[i].Count++
				isExist = true
				break
			}
		}

		if !isExist {
			productCounts = append(productCounts, productCount{
				ProductID: orderProduct.ProductID,
				Count:     1,
			})
		}
	}

	for i := range productCounts {
		for j := range productCounts {
			if productCounts[i].Count > productCounts[j].Count {
				productCounts[i], productCounts[j] = productCounts[j], productCounts[i]
			}
		}
	}

	//get only 20 product
	if len(productCounts) > 20 {
		productCounts = productCounts[:20]
	}

	var recommenddedProducts []httpCommon.Product

	for _, productCount := range productCounts {
		product, err := p.ProductRepository.GetProductByID(ctx, tx, productCount.ProductID)

		if err != nil {
			return nil, err
		}

		recommenddedProducts = append(recommenddedProducts, httpCommon.Product{
			ID:          product.ProductID,
			SellerID:    product.SellerID,
			Name:        product.Name,
			Price:       product.Price,
			Stock:       product.Stock,
			Category:    product.Category,
			Description: product.Description,
			Weight:      product.Weight,
		})
	}

	return recommenddedProducts, nil
}

func (p ProductUsecaseImplementation) GetSearchedProduct(ctx *gin.Context, search string) ([]httpCommon.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) GetAllProductBySellerID(ctx *gin.Context, sellerID string) ([]httpCommon.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) GetProductByID(ctx *gin.Context, productID string) (httpCommon.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) GetOrderByID(ctx *gin.Context, orderID string) (httpCommon.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) GetCustomerOrderByID(ctx *gin.Context, customerID string) ([]httpCommon.OrderItems, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) GetSellerOrderByID(ctx *gin.Context, sellerID string) ([]httpCommon.OrderItems, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) GetOrderProductByID(ctx *gin.Context, orderProductID string) (httpCommon.OrderProduct, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) GetCustomerCart(ctx *gin.Context, customerID string) (httpCommon.CartItems, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) GetPaymentByID(ctx *gin.Context, paymentID string) (httpCommon.Payment, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) InsertProduct(ctx *gin.Context, product httpCommon.AddProduct) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) InsertOrder(ctx *gin.Context, createOrder httpCommon.CreateOrder) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) InsertOrderProduct(ctx *gin.Context, orderProduct httpCommon.OrderProduct) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) InsertCartProduct(ctx *gin.Context, cartProduct httpCommon.CartProduct) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) InsertPayment(ctx *gin.Context, payment httpCommon.Payment) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) UpdateProductByID(ctx *gin.Context, productID string, product httpCommon.Product) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) UpdateOrderStatusByID(ctx *gin.Context, orderID, status string) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) UpdateOrderProductQtyByID(ctx *gin.Context, orderProductID, quantity string) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) UpdateCartProductQtyByID(ctx *gin.Context, cartProductID, quantity string) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) DeleteProductByID(ctx *gin.Context, productID string) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) DeleteOrderProductByID(ctx *gin.Context, orderProductID string) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) DeleteCartProductByID(ctx *gin.Context, cartProductID string) error {
	//TODO implement me
	panic("implement me")
}
