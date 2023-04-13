package product

import (
	"context"
	"database/sql"

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

func (p ProductUsecaseImplementation) GetRecommendedProduct(ctx context.Context) ([]httpCommon.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) GetSearchedProduct(ctx context.Context, search string) ([]httpCommon.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) GetAllProductBySellerID(ctx context.Context, sellerID string) ([]httpCommon.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) GetProductByID(ctx context.Context, productID string) (httpCommon.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) GetOrderByID(ctx context.Context, orderID string) (httpCommon.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) GetCustomerOrderByID(ctx context.Context, customerID string) ([]httpCommon.OrderItems, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) GetSellerOrderByID(ctx context.Context, sellerID string) ([]httpCommon.OrderItems, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) GetOrderProductByID(ctx context.Context, orderProductID string) (httpCommon.OrderProduct, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) GetCustomerCart(ctx context.Context, customerID string) (httpCommon.CartItems, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) GetPaymentByID(ctx context.Context, paymentID string) (httpCommon.Payment, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) InsertProduct(ctx context.Context, product httpCommon.AddProduct) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) InsertOrder(ctx context.Context, createOrder httpCommon.CreateOrder) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) InsertOrderProduct(ctx context.Context, orderProduct httpCommon.OrderProduct) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) InsertCartProduct(ctx context.Context, cartProduct httpCommon.CartProduct) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) InsertPayment(ctx context.Context, payment httpCommon.Payment) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) UpdateProductByID(ctx context.Context, productID string, product httpCommon.Product) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) UpdateOrderStatusByID(ctx context.Context, orderID, status string) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) UpdateOrderProductQtyByID(ctx context.Context, orderProductID, quantity string) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) UpdateCartProductQtyByID(ctx context.Context, cartProductID, quantity string) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) DeleteProductByID(ctx context.Context, productID string) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) DeleteOrderProductByID(ctx context.Context, orderProductID string) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductUsecaseImplementation) DeleteCartProductByID(ctx context.Context, cartProductID string) error {
	//TODO implement me
	panic("implement me")
}
