package product

import (
	"database/sql"

	productRepository "github.com/aziemp66/byte-bargain/internal/repository/product"

	httpCommon "github.com/aziemp66/byte-bargain/common/http"

	"context"
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

func (p *ProductUsecaseImplementation) GetRecommendedProduct(ctx context.Context) error {
	return nil
}

func (p *ProductUsecaseImplementation) GetSearchedProduct(ctx context.Context, search string) (httpCommon.Product, error) {
	return httpCommon.Product{}, nil
}

func (p *ProductUsecaseImplementation) GetAllProductBySellerID(ctx context.Context, sellerID string) ([]httpCommon.Product, error) {
	return nil, nil
}

func (p *ProductUsecaseImplementation) GetProductByID(ctx context.Context, productID string) (httpCommon.Product, error) {
	return httpCommon.Product{}, nil
}

func (p *ProductUsecaseImplementation) GetOrderByID(ctx context.Context, orderID string) error {
	return nil
}

func (p *ProductUsecaseImplementation) GetCustomerOrderByID(ctx context.Context, customerID string) error {
	return nil
}

func (p *ProductUsecaseImplementation) GetSellerOrderByID(ctx context.Context, sellerID string) error {
	return nil
}

func (p *ProductUsecaseImplementation) GetOrderProductByID(ctx context.Context, orderProductID string) error {
	return nil
}

func (p *ProductUsecaseImplementation) GetCustomerCart(ctx context.Context, customerID string) error {
	return nil
}

func (p *ProductUsecaseImplementation) GetPaymentByID(ctx context.Context, paymentID string) error {
	return nil
}

func (p *ProductUsecaseImplementation) InsertProduct(ctx context.Context, product httpCommon.Product) error {
	return nil
}

func (p *ProductUsecaseImplementation) InsertOrder(ctx context.Context, createOrder httpCommon.CreateOrder) error {
	return nil
}

func (p *ProductUsecaseImplementation) InsertOrderProduct(ctx context.Context, orderProduct httpCommon.OrderProduct) error {
	return nil
}

func (p *ProductUsecaseImplementation) InsertCartProduct(ctx context.Context, cartProduct httpCommon.CartProduct) error {
	return nil
}

func (p *ProductUsecaseImplementation) InsertPayment(ctx context.Context, payment httpCommon.Payment) error {
	return nil
}

func (p *ProductUsecaseImplementation) UpdateProductByID(ctx context.Context, productID string, product httpCommon.Product) error {
	return nil
}

func (p *ProductUsecaseImplementation) UpdateOrderStatusByID(ctx context.Context, orderID, status string) error {
	return nil
}

func (p *ProductUsecaseImplementation) UpdateOrderProductQtyByID(ctx context.Context, orderProductID, quantity string) error {
	return nil
}

func (p *ProductUsecaseImplementation) UpdateCartProductQtyByID(ctx context.Context, cartProductID, quantity string) error {
	return nil
}

func (p *ProductUsecaseImplementation) DeleteProductByID(ctx context.Context, productID string) error {
	return nil
}

func (p *ProductUsecaseImplementation) DeleteOrderProductByID(ctx context.Context, orderProductID string) error {
	return nil
}
