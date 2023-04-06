package order

import (
	"context"
	"database/sql"

	errorCommon "github.com/aziemp66/byte-bargain/common/error"
	productDomain "github.com/aziemp66/byte-bargain/internal/domain/product"
)

type ProductRepositoryImplementation struct {
}

func NewProductRepositoryImplementation() *ProductRepositoryImplementation {
	return &ProductRepositoryImplementation{}
}

func (p ProductRepositoryImplementation) GetRecommendedProduct(ctx context.Context, tx *sql.Tx) ([]productDomain.Product, error) {
	//TODO implement me
	panic("implement me")
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
		return product, errorCommon.NewInvariantError(err.Error())
	}

	return product, nil
}

func (p ProductRepositoryImplementation) GetOrderByID(ctx context.Context, tx *sql.Tx, orderID string) (productDomain.Order, error) {
	var order productDomain.Order

	query := `SELECT order_id, customer_id, seller_id, order_date, status FROM orders WHERE order_id = ?`

	err := tx.QueryRowContext(ctx, query, orderID).Scan(&order.OrderID, &order.CustomerID, &order.SellerID, &order.OrderDate, &order.Status)

	if err != nil {
		return order, errorCommon.NewInvariantError(err.Error())
	}

	return order, nil
}

func (p ProductRepositoryImplementation) GetOrderProductByID(ctx context.Context, tx *sql.Tx, orderProductID string) (productDomain.OrderProduct, error) {
	var orderProduct productDomain.OrderProduct

	query := `SELECT order_product_id, order_id, product_id, quantity FROM order_product WHERE order_product_id = ?`

	err := tx.QueryRowContext(ctx, query, orderProductID).Scan(&orderProduct.OrderProductID, &orderProduct.OrderID, &orderProduct.ProductID, &orderProduct.Quantity)

	if err != nil {
		return orderProduct, errorCommon.NewInvariantError(err.Error())
	}

	return orderProduct, nil
}

func (p ProductRepositoryImplementation) GetCartByCustomerID(ctx context.Context, tx *sql.Tx, customerID string) (productDomain.Cart, error) {
	var cart productDomain.Cart

	query := `SELECT cart_id, customer_id FROM cart WHERE customer_id = ?`

	err := tx.QueryRowContext(ctx, query, customerID).Scan(&cart.CartID, &cart.CustomerID)

	if err != nil {
		return cart, errorCommon.NewInvariantError(err.Error())
	}

	return cart, nil
}

func (p ProductRepositoryImplementation) GetCartProductByCartID(ctx context.Context, tx *sql.Tx, cartID string) ([]productDomain.CartProduct, error) {
	var cartProducts []productDomain.CartProduct

	query := `SELECT cart_product_id, cart_id, product_id, quantity FROM cart_product WHERE cart_id = ?`

	rows, err := tx.QueryContext(ctx, query, cartID)

	if err != nil {
		return cartProducts, errorCommon.NewInvariantError(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var cartProduct productDomain.CartProduct

		err = rows.Scan(&cartProduct.CartProductID, &cartProduct.CartID, &cartProduct.ProductID, &cartProduct.Quantity)

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
		return payment, errorCommon.NewInvariantError(err.Error())
	}

	return payment, nil

}

func (p ProductRepositoryImplementation) InsertProduct(ctx context.Context, tx *sql.Tx, sellerID, productName, price, stock, category, description, weight string) error {
	query := `INSERT INTO product (seller_id, name, price, stock, category, description, weight) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := tx.ExecContext(ctx, query, sellerID, productName, price, stock, category, description, weight)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}

func (p ProductRepositoryImplementation) InsertOrder(ctx context.Context, tx *sql.Tx, customerID, sellerID, orderDate, status string) error {
	query := `INSERT INTO orders (customer_id, seller_id, order_date, status) VALUES (?, ?, ?, ?)`

	_, err := tx.ExecContext(ctx, query, customerID, sellerID, orderDate, status)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}

func (p ProductRepositoryImplementation) InsertOrderProduct(ctx context.Context, tx *sql.Tx, orderID, productID, quantity string) error {
	query := `INSERT INTO order_product (order_id, product_id, quantity) VALUES (?, ?, ?)`

	_, err := tx.ExecContext(ctx, query, orderID, productID, quantity)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}

func (p ProductRepositoryImplementation) InsertCartProduct(ctx context.Context, tx *sql.Tx, cartID, productID, quantity string) error {
	query := `INSERT INTO cart_product (cart_id, product_id, quantity) VALUES (?, ?, ?)`

	_, err := tx.ExecContext(ctx, query, cartID, productID, quantity)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}

func (p ProductRepositoryImplementation) InsertPayment(ctx context.Context, tx *sql.Tx, orderID, paymentDate, totalPayment, paymentMethod string) error {
	query := `INSERT INTO payment (order_id, payment_date, total_payment, payment_method) VALUES (?, ?, ?, ?)`

	_, err := tx.ExecContext(ctx, query, orderID, paymentDate, totalPayment, paymentMethod)

	if err != nil {
		return errorCommon.NewInvariantError(err.Error())
	}

	return nil
}

func (p ProductRepositoryImplementation) UpdateProductByID(ctx context.Context, tx *sql.Tx, productID, sellerID, productName, price, stock, category, description, weight string) error {
	query := `UPDATE product SET seller_id = ?, name = ?, price = ?, stock = ?, category = ?, description = ?, weight = ? WHERE product_id = ?`

	_, err := tx.ExecContext(ctx, query, sellerID, productName, price, stock, category, description, weight, productID)

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
	query := `UPDATE order_product SET quantity = ? WHERE order_product_id = ?`

	_, err := tx.ExecContext(ctx, query, quantity, orderProductID)

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
