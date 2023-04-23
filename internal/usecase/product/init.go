package product

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	dbCommon "github.com/aziemp66/byte-bargain/common/db"
	errorCommon "github.com/aziemp66/byte-bargain/common/error"
	httpCommon "github.com/aziemp66/byte-bargain/common/http"
	sessionCommon "github.com/aziemp66/byte-bargain/common/session"

	productRepository "github.com/aziemp66/byte-bargain/internal/repository/product"
)

type ProductUsecaseImplementation struct {
	ProductRepository productRepository.Repository
	DB                *sql.DB
	SessionManager    *sessionCommon.SessionManager
}

func NewProductUsecaseImplementation(productRepository productRepository.Repository, db *sql.DB, sessionManager *sessionCommon.SessionManager) *ProductUsecaseImplementation {
	return &ProductUsecaseImplementation{
		ProductRepository: productRepository,
		DB:                db,
		SessionManager:    sessionManager,
	}
}

func (p *ProductUsecaseImplementation) GetRecommendedProduct(ctx *gin.Context) ([]httpCommon.Product, error) {
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

func (p *ProductUsecaseImplementation) GetSearchedProduct(ctx *gin.Context, search string) ([]httpCommon.Product, error) {
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
		searchedProducts = append(searchedProducts, httpCommon.Product{
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

	return searchedProducts, nil
}

func (p *ProductUsecaseImplementation) GetProductBySellerID(ctx *gin.Context, sellerID string) ([]httpCommon.Product, error) {
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
		sellerProducts = append(sellerProducts, httpCommon.Product{
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

	return sellerProducts, nil
}

func (p *ProductUsecaseImplementation) GetProductByID(ctx *gin.Context, productID string) (httpCommon.Product, error) {
	tx, err := p.DB.Begin()

	if err != nil {
		return httpCommon.Product{}, err
	}

	defer dbCommon.CommitOrRollback(tx)

	product, err := p.ProductRepository.GetProductByID(ctx, tx, productID)

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
	}, nil
}

func (p *ProductUsecaseImplementation) GetOrderByID(ctx *gin.Context, orderID string) (httpCommon.Order, error) {
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

func (p *ProductUsecaseImplementation) GetOrderByCustomerID(ctx *gin.Context, customerID string) ([]httpCommon.OrderItems, error) {
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

func (p *ProductUsecaseImplementation) GetSellerOrderByID(ctx *gin.Context, sellerID string) ([]httpCommon.OrderItems, error) {
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

func (p *ProductUsecaseImplementation) GetOrderProductByID(ctx *gin.Context, orderProductID string) (httpCommon.OrderProduct, error) {
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

func (p *ProductUsecaseImplementation) GetCustomerCart(ctx *gin.Context, customerID string) (httpCommon.Cart, error) {
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
			Price:         product.Price,
			Qty:           v.Quantity,
		})
	}

	return httpCommon.Cart{
		Items:        products,
		TotalPayment: totalPayment,
	}, nil
}

func (p *ProductUsecaseImplementation) GetPaymentByID(ctx *gin.Context, paymentID string) (httpCommon.Payment, error) {
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

func (p *ProductUsecaseImplementation) InsertProduct(ctx *gin.Context, product httpCommon.AddProduct) error {
	tx, err := p.DB.Begin()

	if err != nil {
		return err
	}

	defer dbCommon.CommitOrRollback(tx)

	sellerID := p.SessionManager.GetSessionValue(ctx, "seller_id")

	if sellerID == "" {
		return errorCommon.NewForbiddenError("You are not a seller")
	}

	sellerIDString, ok := sellerID.(string)

	if !ok {
		return errorCommon.NewInvariantError("seller_id is not string")
	}

	err = p.ProductRepository.InsertProduct(ctx, tx, sellerIDString, product.Name, product.Price, product.Stock, product.Category, product.Description, product.Weight)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductUsecaseImplementation) InsertOrder(ctx *gin.Context, createOrder httpCommon.CreateOrder) error {
	tx, err := p.DB.Begin()

	if err != nil {
		return err
	}

	defer dbCommon.CommitOrRollback(tx)

	customerID := p.SessionManager.GetSessionValue(ctx, "customer_id")

	if customerID == "" {
		return errorCommon.NewForbiddenError("You are not a customer")
	}

	customerIDString, ok := customerID.(string)

	if !ok {
		return errorCommon.NewInvariantError("customer_id is not string")
	}

	err = p.ProductRepository.InsertOrder(ctx, tx, customerIDString, createOrder.SellerID, time.Now(), "pending")

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductUsecaseImplementation) InsertOrderProduct(ctx *gin.Context, orderProduct httpCommon.OrderProduct) error {
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

func (p *ProductUsecaseImplementation) InsertCartProduct(ctx *gin.Context, cartProduct httpCommon.CartProduct) error {
	tx, err := p.DB.Begin()

	if err != nil {
		return err
	}

	defer dbCommon.CommitOrRollback(tx)

	customerID := p.SessionManager.GetSessionValue(ctx, "customer_id")

	if customerID == "" {
		return errorCommon.NewForbiddenError("You are not a customer")
	}

	customerIDString, ok := customerID.(string)

	if !ok {
		return errorCommon.NewInvariantError("customer_id is not string")
	}

	err = p.ProductRepository.InsertCartProduct(ctx, tx, customerIDString, cartProduct.ProductID, cartProduct.Qty)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductUsecaseImplementation) InsertPayment(ctx *gin.Context, payment httpCommon.Payment) error {
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

func (p *ProductUsecaseImplementation) InsertImages(ctx *gin.Context) (imagesID []string, err error) {
	tx, err := p.DB.Begin()

	if err != nil {
		return nil, err
	}

	defer dbCommon.CommitOrRollback(tx)

	form, err := ctx.MultipartForm()

	if err != nil {
		return nil, errorCommon.NewInvariantError("Error reading multipart form")
	}

	files := form.File["images"]

	for _, file := range files {
		fileName := fmt.Sprintf("%s_%s", uuid.New().String(), file.Filename)

		filePath := "public/product_images/" + fileName

		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			return nil, errorCommon.NewInvariantError("Error saving file")
		}

		imageID, err := p.ProductRepository.InsertImage(ctx, tx, filePath)

		if err != nil {
			return nil, err
		}

		imagesID = append(imagesID, imageID)
	}

	return imagesID, nil
}

func (p *ProductUsecaseImplementation) UpdateProductByID(ctx *gin.Context, productID string, product httpCommon.Product) error {
	//TODO implement me
	panic("implement me")
}

func (p *ProductUsecaseImplementation) UpdateOrderStatusByID(ctx *gin.Context, orderID, status string) error {
	//TODO implement me
	panic("implement me")
}

func (p *ProductUsecaseImplementation) UpdateOrderProductQtyByID(ctx *gin.Context, orderProductID, quantity string) error {
	//TODO implement me
	panic("implement me")
}

func (p *ProductUsecaseImplementation) UpdateCartProductQtyByID(ctx *gin.Context, cartProductID, quantity string) error {
	//TODO implement me
	panic("implement me")
}

func (p *ProductUsecaseImplementation) DeleteProductByID(ctx *gin.Context, productID string) error {
	//TODO implement me
	panic("implement me")
}

func (p *ProductUsecaseImplementation) DeleteOrderProductByID(ctx *gin.Context, orderProductID string) error {
	//TODO implement me
	panic("implement me")
}

func (p *ProductUsecaseImplementation) DeleteCartProductByID(ctx *gin.Context, cartProductID string) error {
	//TODO implement me
	panic("implement me")
}
