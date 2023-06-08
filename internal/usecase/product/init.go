package product

import (
	"context"
	"database/sql"
	"time"

	dbCommon "github.com/aziemp66/byte-bargain/common/db"
	errorCommon "github.com/aziemp66/byte-bargain/common/error"
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

func (p *ProductUsecaseImplementation) GetRecommendedProduct(ctx context.Context) ([]httpCommon.Product, error) {
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

		image, err := p.ProductRepository.GetImageByProductID(ctx, tx, product.ProductID)

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
			Image:       "/product_image/" + image.Image,
		})
	}

	return recommenddedProducts, nil
}

func (p *ProductUsecaseImplementation) GetSearchedProduct(ctx context.Context, search string) ([]httpCommon.Product, error) {
	tx, err := p.DB.Begin()

	if err != nil {
		return nil, err
	}

	defer dbCommon.CommitOrRollback(tx)

	products, err := p.ProductRepository.GetSearchedProduct(ctx, tx, search)

	if err != nil {
		return nil, err
	}

	var searchedProducts []httpCommon.Product

	for _, product := range products {
		image, err := p.ProductRepository.GetImageByProductID(ctx, tx, product.ProductID)

		if err != nil {
			return nil, err
		}
		searchedProducts = append(searchedProducts, httpCommon.Product{
			ID:          product.ProductID,
			SellerID:    product.SellerID,
			Name:        product.Name,
			Price:       product.Price,
			Stock:       product.Stock,
			Category:    product.Category,
			Description: product.Description,
			Weight:      product.Weight,
			Image:       "/product_image/" + image.Image,
		})
	}

	return searchedProducts, nil
}

func (p *ProductUsecaseImplementation) GetProductBySellerID(ctx context.Context, sellerID string) ([]httpCommon.Product, error) {
	tx, err := p.DB.Begin()

	if err != nil {
		return nil, err
	}

	defer dbCommon.CommitOrRollback(tx)

	products, err := p.ProductRepository.GetAllProductBySellerID(ctx, tx, sellerID)

	if err != nil {
		return nil, err
	}

	var sellerProducts []httpCommon.Product

	for _, product := range products {
		image, err := p.ProductRepository.GetImageByProductID(ctx, tx, product.ProductID)

		if err != nil {
			return nil, err
		}

		sellerProducts = append(sellerProducts, httpCommon.Product{
			ID:          product.ProductID,
			SellerID:    product.SellerID,
			Name:        product.Name,
			Price:       product.Price,
			Stock:       product.Stock,
			Category:    product.Category,
			Description: product.Description,
			Weight:      product.Weight,
			Image:       "/product_image/" + image.Image,
		})
	}

	return sellerProducts, nil
}

func (p *ProductUsecaseImplementation) GetProductByID(ctx context.Context, productID string) (httpCommon.Product, error) {
	tx, err := p.DB.Begin()

	if err != nil {
		return httpCommon.Product{}, err
	}

	defer dbCommon.CommitOrRollback(tx)

	product, err := p.ProductRepository.GetProductByID(ctx, tx, productID)

	if err != nil {
		return httpCommon.Product{}, err
	}

	image, err := p.ProductRepository.GetImageByProductID(ctx, tx, product.ProductID)

	if err != nil {
		return httpCommon.Product{}, err
	}

	return httpCommon.Product{
		ID:          product.ProductID,
		SellerID:    product.SellerID,
		Name:        product.Name,
		Price:       product.Price,
		Stock:       product.Stock,
		Category:    product.Category,
		Description: product.Description,
		Weight:      product.Weight,
		Image:       "/product_image/" + image.Image,
	}, nil
}

func (p *ProductUsecaseImplementation) GetOrderByID(ctx context.Context, orderID string) (httpCommon.Order, error) {
	tx, err := p.DB.Begin()

	if err != nil {
		return httpCommon.Order{}, err
	}

	defer dbCommon.CommitOrRollback(tx)

	order, err := p.ProductRepository.GetOrderByID(ctx, tx, orderID)

	if err != nil {
		return httpCommon.Order{}, err
	}

	return httpCommon.Order{
		OrderID:    order.OrderID,
		CustomerID: order.CustomerID,
		SellerID:   order.SellerID,
		OrderDate:  order.OrderDate,
		Status:     order.Status,
	}, nil
}

func (p *ProductUsecaseImplementation) GetOrderByCustomerID(ctx context.Context, customerID string) ([]httpCommon.OrderItems, error) {
	tx, err := p.DB.Begin()

	if err != nil {
		return nil, err
	}

	defer dbCommon.CommitOrRollback(tx)

	orders, err := p.ProductRepository.GetOrderByCustomerID(ctx, tx, customerID)

	if err != nil {
		return nil, err
	}

	var orderItems []httpCommon.OrderItems

	for _, v := range orders {
		orderProduct, err := p.ProductRepository.GetOrderProductByOrderID(ctx, tx, v.OrderID)

		if err != nil {
			return nil, err
		}

		var products []httpCommon.ProductItem
		var totalPayment float64

		for _, v2 := range orderProduct {
			products = append(products, httpCommon.ProductItem{
				ProductID: v2.ProductID,
				Qty:       v2.Quantity,
			})

			product, err := p.ProductRepository.GetProductByID(ctx, tx, v2.ProductID)

			if err != nil {
				return nil, err
			}

			totalPayment += product.Price * float64(v2.Quantity)
		}

		orderItems = append(orderItems, httpCommon.OrderItems{
			OrderID:      v.OrderID,
			CustomerID:   v.CustomerID,
			SellerID:     v.SellerID,
			OrderDate:    v.OrderDate,
			Status:       v.Status,
			Products:     products,
			TotalPayment: totalPayment,
		})
	}

	return orderItems, nil
}

func (p *ProductUsecaseImplementation) GetOrderBySellerID(ctx context.Context, sellerID string) ([]httpCommon.OrderItems, error) {
	tx, err := p.DB.Begin()

	if err != nil {
		return nil, err
	}

	defer dbCommon.CommitOrRollback(tx)

	orders, err := p.ProductRepository.GetOrderBySellerID(ctx, tx, sellerID)

	if err != nil {
		return nil, err
	}

	var orderItems []httpCommon.OrderItems

	for _, v := range orders {
		orderProduct, err := p.ProductRepository.GetOrderProductByOrderID(ctx, tx, v.OrderID)

		if err != nil {
			return nil, err
		}

		var products []httpCommon.ProductItem
		var totalPayment float64

		for _, v2 := range orderProduct {
			products = append(products, httpCommon.ProductItem{
				ProductID: v2.ProductID,
				Qty:       v2.Quantity,
			})

			product, err := p.ProductRepository.GetProductByID(ctx, tx, v2.ProductID)

			if err != nil {
				return nil, err
			}

			totalPayment += product.Price * float64(v2.Quantity)
		}

		orderItems = append(orderItems, httpCommon.OrderItems{
			OrderID:      v.OrderID,
			CustomerID:   v.CustomerID,
			SellerID:     v.SellerID,
			OrderDate:    v.OrderDate,
			Status:       v.Status,
			Products:     products,
			TotalPayment: totalPayment,
		})
	}

	return orderItems, nil
}

func (p *ProductUsecaseImplementation) GetCartProductByID(ctx context.Context, cartProductID string) (httpCommon.CartProduct, error) {
	tx, err := p.DB.Begin()

	if err != nil {
		return httpCommon.CartProduct{}, err
	}

	defer dbCommon.CommitOrRollback(tx)

	cartProduct, err := p.ProductRepository.GetCartProductByID(ctx, tx, cartProductID)

	if err != nil {
		return httpCommon.CartProduct{}, err
	}

	product, err := p.ProductRepository.GetProductByID(ctx, tx, cartProduct.ProductID)

	if err != nil {
		return httpCommon.CartProduct{}, err
	}

	return httpCommon.CartProduct{
		CartProductID: cartProduct.CartProductID,
		ProductID:     cartProduct.ProductID,
		CustomerID:    cartProduct.CustomerID,
		Price:         product.Price,
		Qty:           cartProduct.Quantity,
	}, nil
}

func (p *ProductUsecaseImplementation) GetOrderProductByID(ctx context.Context, orderProductID string) (httpCommon.OrderProduct, error) {
	tx, err := p.DB.Begin()

	if err != nil {
		return httpCommon.OrderProduct{}, err
	}

	defer dbCommon.CommitOrRollback(tx)

	orderProduct, err := p.ProductRepository.GetOrderProductByID(ctx, tx, orderProductID)

	if err != nil {
		return httpCommon.OrderProduct{}, err
	}

	return httpCommon.OrderProduct{
		OrderProductID: orderProduct.OrderProductID,
		OrderID:        orderProduct.OrderID,
		ProductID:      orderProduct.ProductID,
		Qty:            orderProduct.Quantity,
	}, nil
}

func (p *ProductUsecaseImplementation) GetCustomerCart(ctx context.Context, customerID string) (httpCommon.Cart, error) {
	tx, err := p.DB.Begin()

	if err != nil {
		return httpCommon.Cart{}, err
	}

	defer dbCommon.CommitOrRollback(tx)

	cartProduct, err := p.ProductRepository.GetCartProductByCustomerID(ctx, tx, customerID)

	if err != nil {
		return httpCommon.Cart{}, err
	}

	var products []httpCommon.CartProduct
	var totalPayment float64

	for _, v := range cartProduct {
		product, err := p.ProductRepository.GetProductByID(ctx, tx, v.ProductID)

		if err != nil {
			return httpCommon.Cart{}, err
		}

		totalPayment += product.Price * float64(v.Quantity)

		products = append(products, httpCommon.CartProduct{
			CartProductID: v.CartProductID,
			ProductID:     v.ProductID,
			Name:          product.Name,
			Price:         product.Price,
			Qty:           v.Quantity,
		})
	}

	return httpCommon.Cart{
		Items:        products,
		TotalPayment: totalPayment,
	}, nil
}

func (p *ProductUsecaseImplementation) GetPaymentByID(ctx context.Context, paymentID string) (httpCommon.Payment, error) {
	tx, err := p.DB.Begin()

	if err != nil {
		return httpCommon.Payment{}, err
	}

	defer dbCommon.CommitOrRollback(tx)

	payment, err := p.ProductRepository.GetPaymentByID(ctx, tx, paymentID)

	if err != nil {
		return httpCommon.Payment{}, err
	}

	return httpCommon.Payment{
		PaymentID:     payment.PaymentID,
		OrderID:       payment.OrderID,
		PaymentDate:   payment.PaymentDate.Format("2006-01-02 15:04:05"),
		TotalPayment:  payment.TotalPayment,
		PaymentMethod: payment.PaymentMethod,
	}, nil
}

func (p *ProductUsecaseImplementation) InsertProduct(ctx context.Context, sellerID string, product httpCommon.ProductRequest) error {
	tx, err := p.DB.Begin()

	if err != nil {
		return err
	}

	defer dbCommon.CommitOrRollback(tx)

	productID, err := p.ProductRepository.InsertProduct(ctx, tx, sellerID, product.Name, product.Price, product.Stock, product.Category, product.Description, product.Weight)

	if err != nil {
		return err
	}

	err = p.ProductRepository.UpdateLinkImageByID(ctx, tx, product.Image, productID)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductUsecaseImplementation) InsertOrder(ctx context.Context, customerID string, createOrder httpCommon.CreateOrder) error {
	tx, err := p.DB.Begin()

	if err != nil {
		return err
	}

	defer dbCommon.CommitOrRollback(tx)

	err = p.ProductRepository.InsertOrder(ctx, tx, customerID, createOrder.SellerID, time.Now(), "pending")

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductUsecaseImplementation) InsertOrderProduct(ctx context.Context, orderProduct httpCommon.OrderProduct) error {
	tx, err := p.DB.Begin()

	if err != nil {
		return err
	}

	defer dbCommon.CommitOrRollback(tx)

	err = p.ProductRepository.InsertOrderProduct(ctx, tx, orderProduct.OrderID, orderProduct.ProductID, orderProduct.Qty)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductUsecaseImplementation) InsertCartProduct(ctx context.Context, customerID string, cartProduct httpCommon.AddCartProduct) error {
	tx, err := p.DB.Begin()

	if err != nil {
		return err
	}

	defer dbCommon.CommitOrRollback(tx)

	cartProducts, err := p.ProductRepository.GetCartProductByCustomerID(ctx, tx, customerID)

	if err != nil {
		return err
	}

	for _, v := range cartProducts {
		if v.ProductID == cartProduct.ProductID {
			err = p.ProductRepository.UpdateCartProductQtyByID(ctx, tx, v.CartProductID, v.Quantity+cartProduct.Qty)

			if err != nil {
				return err
			}

			return nil
		}

	}

	err = p.ProductRepository.InsertCartProduct(ctx, tx, customerID, cartProduct.ProductID, cartProduct.Qty)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductUsecaseImplementation) InsertPayment(ctx context.Context, payment httpCommon.Payment) error {
	tx, err := p.DB.Begin()

	if err != nil {
		return err
	}

	defer dbCommon.CommitOrRollback(tx)

	err = p.ProductRepository.InsertPayment(ctx, tx, payment.OrderID, time.Now(), payment.TotalPayment, payment.PaymentMethod)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductUsecaseImplementation) InsertImage(ctx context.Context, filePath string) (imagesID string, err error) {
	tx, err := p.DB.Begin()

	if err != nil {
		return "", err
	}

	defer dbCommon.CommitOrRollback(tx)

	imagesID, err = p.ProductRepository.InsertImage(ctx, tx, filePath)

	return imagesID, nil
}

func (p *ProductUsecaseImplementation) UpdateProductByID(ctx context.Context, productID string, product httpCommon.ProductRequest) error {
	tx, err := p.DB.Begin()

	if err != nil {
		return err
	}

	defer dbCommon.CommitOrRollback(tx)

	productRes, err := p.ProductRepository.GetProductByID(ctx, tx, productID)

	if err != nil {
		return err
	}

	if productRes.ProductID == "" {
		return errorCommon.NewNotFoundError("Product not found")
	}

	err = p.ProductRepository.UpdateProductByID(ctx, tx, productID, product.Name, product.Price, product.Stock, product.Category, product.Description, product.Weight)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductUsecaseImplementation) UpdateOrderStatusByID(ctx context.Context, updateOrderStatus httpCommon.UpdateOrderStatus) error {
	tx, err := p.DB.Begin()

	if err != nil {
		return err
	}

	defer dbCommon.CommitOrRollback(tx)

	if updateOrderStatus.Status != httpCommon.ORDER_STATUS_CANCELED &&
		updateOrderStatus.Status != httpCommon.ORDER_STATUS_DELIVERED &&
		updateOrderStatus.Status != httpCommon.ORDER_STATUS_SHIPPED &&
		updateOrderStatus.Status != httpCommon.ORDER_STATUS_PROCESSING &&
		updateOrderStatus.Status != httpCommon.ORDER_STATUS_PENDING &&
		updateOrderStatus.Status != httpCommon.ORDER_STATUS_REJECTED &&
		updateOrderStatus.Status != httpCommon.ORDER_STATUS_REFUNDED &&
		updateOrderStatus.Status != httpCommon.ORDER_STATUS_FAILED &&
		updateOrderStatus.Status != httpCommon.ORDER_STATUS_COMPLETED &&
		updateOrderStatus.Status != httpCommon.ORDER_STATUS_RETURNED {
		return errorCommon.NewInvariantError("Invalid order status")
	}

	err = p.ProductRepository.UpdateOrderStatusByID(ctx, tx, updateOrderStatus.OrderID, updateOrderStatus.Status)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductUsecaseImplementation) UpdateOrderProductQtyByID(ctx context.Context, orderProductID string, quantity int) error {
	tx, err := p.DB.Begin()

	if err != nil {
		return err
	}

	defer dbCommon.CommitOrRollback(tx)

	err = p.ProductRepository.UpdateOrderProductQtyByID(ctx, tx, orderProductID, quantity)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductUsecaseImplementation) UpdateCartProductQtyByID(ctx context.Context, cartProductID string, quantity int) error {
	tx, err := p.DB.Begin()

	if err != nil {
		return err
	}

	defer dbCommon.CommitOrRollback(tx)

	err = p.ProductRepository.UpdateCartProductQtyByID(ctx, tx, cartProductID, quantity)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductUsecaseImplementation) DeleteProductByID(ctx context.Context, productID string) error {
	tx, err := p.DB.Begin()

	if err != nil {
		return err
	}

	defer dbCommon.CommitOrRollback(tx)

	err = p.ProductRepository.DeleteProductByID(ctx, tx, productID)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductUsecaseImplementation) DeleteCartProductByID(ctx context.Context, cartProductID string) error {
	tx, err := p.DB.Begin()

	if err != nil {
		return err
	}

	defer dbCommon.CommitOrRollback(tx)

	err = p.ProductRepository.DeleteCartProductByID(ctx, tx, cartProductID)

	if err != nil {
		return err
	}

	return nil
}
