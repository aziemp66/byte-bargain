package order

import (
	"context"
	"database/sql"
	"time"

	errorCommon "github.com/aziemp66/byte-bargain/common/error"
	productDomain "github.com/aziemp66/byte-bargain/internal/domain/product"
	"github.com/google/uuid"
)

type ProductRepositoryImplementation struct {
}

func NewProductRepositoryImplementation() *ProductRepositoryImplementation {
	return &ProductRepositoryImplementation{}
}

func (p ProductRepositoryImplementation) GetAllProduct(ctx context.Context, tx *sql.Tx) ([]productDomain.Product, error) {
	var products []productDomain.Product

	query := `SELECT product_id, seller_id, name, price, stock, category, description, weight FROM product`

	rows, err := tx.QueryContext(ctx, query)

	if err != nil {
		return products, errorCommon.NewInvariantError(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var product productDomain.Product

		err = rows.Scan(&product.ProductID, &product.SellerID, &product.Name, &product.Price, &product.Stock, &product.Category, &product.Description, &product.Weight)

		if err != nil {
			return products, errorCommon.NewInvariantError(err.Error())
		}

		products = append(products, product)
	}

	return products, nil
}

func (p ProductRepositoryImplementation) GetSearchedProduct(ctx context.Context, tx *sql.Tx, search string) ([]productDomain.Product, error) {
	var products []productDomain.Product

	query := `SELECT product_id, seller_id, name, price, stock, category, description, weight FROM product WHERE name LIKE ?`

	rows, err := tx.QueryContext(ctx, query, "%"+search+"%")

	if err != nil {
		return products, errorCommon.NewInvariantError(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var product productDomain.Product

		err = rows.Scan(&product.ProductID, &product.SellerID, &product.Name, &product.Price, &product.Stock, &product.Category, &product.Description, &product.Weight)

		if err != nil {
			return products, errorCommon.NewInvariantError(err.Error())
		}

		products = append(products, product)
	}

	return products, nil
}

func (p ProductRepositoryImplementation) GetAllProductBySellerID(ctx context.Context, tx *sql.Tx, sellerID string) ([]productDomain.Product, error) {
	var products []productDomain.Product

	query := `SELECT product_id, seller_id, name, price, stock, category, description, weight FROM product WHERE seller_id = ?`

	rows, err := tx.QueryContext(ctx, query, sellerID)

	if err != nil {
		return products, errorCommon.NewInvariantError(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var product productDomain.Product

		err = rows.Scan(&product.ProductID, &product.SellerID, &product.Name, &product.Price, &product.Stock, &product.Category, &product.Description, &product.Weight)

		if err != nil {
			return products, errorCommon.NewInvariantError(err.Error())
		}

		products = append(products, product)
	}

	return products, nil
}

func (p ProductRepositoryImplementation) GetProductByID(ctx context.Context, tx *sql.Tx, productID string) (productDomain.Product, error) {
	var product productDomain.Product

	query := `SELECT product_id, seller_id, name, price, stock, category, description, weight FROM product WHERE product_id = ?`

	err := tx.QueryRowContext(ctx, query, productID).Scan(&product.ProductID, &product.SellerID, &product.Name, &product.Price, &product.Stock, &product.Category, &product.Description, &product.Weight)

	if err != nil {
		return product, errorCommon.NewInvariantError("product not found")
	}

	return product, nil
}

func (p ProductRepositoryImplementation) GetOrderByID(ctx context.Context, tx *sql.Tx, orderID string) (productDomain.Order, error) {
	var order productDomain.Order

	query := `SELECT order_id, customer_id, seller_id, order_date, status FROM orders WHERE order_id = ?`

	err := tx.QueryRowContext(ctx, query, orderID).Scan(&order.OrderID, &order.CustomerID, &order.SellerID, &order.OrderDate, &order.Status)

	if err != nil {
		return order, errorCommon.NewInvariantError("order not found")
	}

	return order, nil
}

func (p ProductRepositoryImplementation) GetOrderProductByID(ctx context.Context, tx *sql.Tx, orderProductID string) (productDomain.OrderProduct, error) {
	var orderProduct productDomain.OrderProduct

	query := `SELECT order_product_id, order_id, product_id, quantity FROM order_product WHERE order_product_id = ?`

	err := tx.QueryRowContext(ctx, query, orderProductID).Scan(&orderProduct.OrderProductID, &orderProduct.OrderID, &orderProduct.ProductID, &orderProduct.Quantity)

	if err != nil {
		return orderProduct, errorCommon.NewInvariantError("order product not found")
	}

	return orderProduct, nil
}

func (p ProductRepositoryImplementation) GetOrderByCustomerID(ctx context.Context, tx *sql.Tx, customerID string) ([]productDomain.Order, error) {
	var orders []productDomain.Order

	query := `SELECT order_id, customer_id, seller_id, order_date, status FROM orders WHERE customer_id = ?`

	rows, err := tx.QueryContext(ctx, query, customerID)

	if err != nil {
		return orders, errorCommon.NewInvariantError(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var order productDomain.Order

		err = rows.Scan(&order.OrderID, &order.CustomerID, &order.SellerID, &order.OrderDate, &order.Status)

		if err != nil {
			return orders, errorCommon.NewInvariantError(err.Error())
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (p ProductRepositoryImplementation) GetOrderBySellerID(ctx context.Context, tx *sql.Tx, sellerID string) ([]productDomain.Order, error) {
	var orders []productDomain.Order

	query := `SELECT order_id, customer_id, seller_id, order_date, status FROM orders WHERE seller_id = ?`

	rows, err := tx.QueryContext(ctx, query, sellerID)

	if err != nil {
		return orders, errorCommon.NewInvariantError(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var order productDomain.Order

		err = rows.Scan(&order.OrderID, &order.CustomerID, &order.SellerID, &order.OrderDate, &order.Status)

		if err != nil {
			return orders, errorCommon.NewInvariantError(err.Error())
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (p ProductRepositoryImplementation) GetCartProductByID(ctx context.Context, tx *sql.Tx, cartProductID string) (productDomain.CartProduct, error) {
	var cartProduct productDomain.CartProduct

	query := `SELECT cart_product_id, customer_id, product_id, quantity FROM cart_product WHERE cart_product_id = ?`

	err := tx.QueryRowContext(ctx, query, cartProductID).Scan(&cartProduct.CartProductID, &cartProduct.CustomerID, &cartProduct.ProductID, &cartProduct.Quantity)

	if err != nil {
		return cartProduct, errorCommon.NewInvariantError("cart product not found")
	}

	return cartProduct, nil
}

func (p ProductRepositoryImplementation) GetCartProductByCustomerID(ctx context.Context, tx *sql.Tx, customerID string) ([]productDomain.CartProduct, error) {
	var cartProducts []productDomain.CartProduct

	query := `SELECT cart_product_id, customer_id, product_id, quantity FROM cart_product WHERE customer_id = ?`

	rows, err := tx.QueryContext(ctx, query, customerID)

	if err != nil {
		return cartProducts, errorCommon.NewInvariantError(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var cartProduct productDomain.CartProduct

		err = rows.Scan(&cartProduct.CartProductID, &cartProduct.CustomerID, &cartProduct.ProductID, &cartProduct.Quantity)

		if err != nil {
			return cartProducts, errorCommon.NewInvariantError(err.Error())
		}

		cartProducts = append(cartProducts, cartProduct)
	}

	return cartProducts, nil
}

func (p ProductRepositoryImplementation) GetPaymentByID(ctx context.Context, tx *sql.Tx, paymentID string) (productDomain.Payment, error) {
	var payment productDomain.Payment

	query := `SELECT payment_id, order_id, payment_date, payment_method FROM payment WHERE payment_id = ?`

	err := tx.QueryRowContext(ctx, query, paymentID).Scan(&payment.PaymentID, &payment.OrderID, &payment.PaymentDate, &payment.PaymentMethod)

	if err != nil {
		return payment, errorCommon.NewInvariantError("payment not found")
	}

	return payment, nil

}

func (p ProductRepositoryImplementation) GetOrderProductByOrderID(ctx context.Context, tx *sql.Tx, orderID string) ([]productDomain.OrderProduct, error) {
	var orderProducts []productDomain.OrderProduct

	query := `SELECT order_product_id, order_id, product_id, quantity FROM order_product WHERE order_id = ?`

	rows, err := tx.QueryContext(ctx, query, orderID)

	if err != nil {
		return orderProducts, errorCommon.NewInvariantError(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var orderProduct productDomain.OrderProduct

		err = rows.Scan(&orderProduct.OrderProductID, &orderProduct.OrderID, &orderProduct.ProductID, &orderProduct.Quantity)

		if err != nil {
			return orderProducts, errorCommon.NewInvariantError(err.Error())
		}

		orderProducts = append(orderProducts, orderProduct)
	}

	return orderProducts, nil
}

func (p ProductRepositoryImplementation) GetAllOrderProduct(ctx context.Context, tx *sql.Tx) ([]productDomain.OrderProduct, error) {
	var orderProducts []productDomain.OrderProduct

	query := `SELECT order_product_id, order_id, product_id, quantity FROM order_product`

	rows, err := tx.QueryContext(ctx, query)

	if err != nil {
		return orderProducts, errorCommon.NewInvariantError(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var orderProduct productDomain.OrderProduct

		err = rows.Scan(&orderProduct.OrderProductID, &orderProduct.OrderID, &orderProduct.ProductID, &orderProduct.Quantity)

		if err != nil {
			return orderProducts, errorCommon.NewInvariantError(err.Error())
		}

		orderProducts = append(orderProducts, orderProduct)
	}

	return orderProducts, nil
}

func (p ProductRepositoryImplementation) InsertProduct(ctx context.Context, tx *sql.Tx, sellerID, productName string, price float64, stock int, category, description string, weight float64) error {
	query := `INSERT INTO product (product_id, seller_id, name, price, stock, category, description, weight) VALUES (?, ?, ?, ?, ?, ?, ?)`

	productID := uuid.New().String()

	_, err := tx.ExecContext(ctx, query, productID, sellerID, productName, price, stock, category, description, weight)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}

func (p ProductRepositoryImplementation) InsertOrder(ctx context.Context, tx *sql.Tx, customerID, sellerID string, orderDate time.Time, status string) error {
	query := `INSERT INTO orders (order_id, customer_id, seller_id, order_date, status) VALUES (?, ?, ?, ?)`

	orderID := uuid.New().String()

	_, err := tx.ExecContext(ctx, query, orderID, customerID, sellerID, orderDate, status)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}

func (p ProductRepositoryImplementation) InsertOrderProduct(ctx context.Context, tx *sql.Tx, orderID, productID string, quantity int) error {
	query := `INSERT INTO order_product (order_product_id, order_id, product_id, quantity) VALUES (?, ?, ?)`

	orderProductID := uuid.New().String()

	_, err := tx.ExecContext(ctx, query, orderProductID, orderID, productID, quantity)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}

func (p ProductRepositoryImplementation) InsertCartProduct(ctx context.Context, tx *sql.Tx, customerID, productID string, quantity int) error {
	query := `INSERT INTO cart_product (cart_product_id, customer_id, product_id, quantity) VALUES (?, ?, ?)`

	cartProductID := uuid.New().String()

	_, err := tx.ExecContext(ctx, query, cartProductID, customerID, productID, quantity)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}

func (p ProductRepositoryImplementation) InsertPayment(ctx context.Context, tx *sql.Tx, orderID string, paymentDate time.Time, totalPayment float64, paymentMethod string) error {
	query := `INSERT INTO payment (payment_id, order_id, payment_date, total_payment, payment_method) VALUES (?, ?, ?, ?)`

	paymentID := uuid.New().String()

	_, err := tx.ExecContext(ctx, query, paymentID, orderID, paymentDate, totalPayment, paymentMethod)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}

func (p ProductRepositoryImplementation) InsertImage(ctx context.Context, tx *sql.Tx, image string) (imageID string, err error) {
	query := `INSERT INTO image (product_image_id, image) VALUES (?, ?)`

	productImageID := uuid.New().String()

	_, err = tx.ExecContext(ctx, query, productImageID, image)

	if err != nil {
		return "", errorCommon.NewInvariantError(err.Error())
	}

	return productImageID, nil
}

func (p ProductRepositoryImplementation) UpdateProductByID(ctx context.Context, tx *sql.Tx, productID, productName string, price float64, stock int, category, description string, weight float64) error {
	query := `UPDATE product SET name = ?, price = ?, stock = ?, category = ?, description = ?, weight = ? WHERE product_id = ?`

	_, err := tx.ExecContext(ctx, query, productName, price, stock, category, description, weight, productID)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}

func (p ProductRepositoryImplementation) UpdateOrderStatusByID(ctx context.Context, tx *sql.Tx, orderID, status string) error {
	query := `UPDATE orders SET status = ? WHERE order_id = ?`

	_, err := tx.ExecContext(ctx, query, status, orderID)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}

func (p ProductRepositoryImplementation) UpdateOrderProductQtyByID(ctx context.Context, tx *sql.Tx, orderProductID, quantity string) error {
	query := `UPDATE order_product SET qty = ? WHERE order_product_id = ?`

	_, err := tx.ExecContext(ctx, query, quantity, orderProductID)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}

func (p ProductRepositoryImplementation) UpdateCartProductQtyByID(ctx context.Context, tx *sql.Tx, cartProductID, quantity string) error {
	query := `UPDATE cart_product SET qty = ? WHERE cart_product_id = ?`

	_, err := tx.ExecContext(ctx, query, quantity, cartProductID)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}

func (p ProductRepositoryImplementation) UpdateLinkImageByID(ctx context.Context, tx *sql.Tx, imageID, productID string) error {
	query := `UPDATE image SET product_id = ? WHERE product_image_id = ?`

	_, err := tx.ExecContext(ctx, query, productID, imageID)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}

func (p ProductRepositoryImplementation) DeleteProductByID(ctx context.Context, tx *sql.Tx, productID string) error {
	query := `DELETE FROM product WHERE product_id = ?`

	_, err := tx.ExecContext(ctx, query, productID)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}

func (p ProductRepositoryImplementation) DeleteOrderProductByID(ctx context.Context, tx *sql.Tx, orderProductID string) error {
	query := `DELETE FROM order_product WHERE order_product_id = ?`

	_, err := tx.ExecContext(ctx, query, orderProductID)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}

func (p ProductRepositoryImplementation) DeleteCartProductByID(ctx context.Context, tx *sql.Tx, cartProductID string) error {
	query := `DELETE FROM cart_product WHERE cart_product_id = ?`

	_, err := tx.ExecContext(ctx, query, cartProductID)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}
