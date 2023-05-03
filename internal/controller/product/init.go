package product

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	errorCommon "github.com/aziemp66/byte-bargain/common/error"
	httpCommon "github.com/aziemp66/byte-bargain/common/http"
	httpMiddleware "github.com/aziemp66/byte-bargain/common/http/middleware"
	sessionCommon "github.com/aziemp66/byte-bargain/common/session"

	productUseCase "github.com/aziemp66/byte-bargain/internal/usecase/product"
	userUsecase "github.com/aziemp66/byte-bargain/internal/usecase/user"
)

type ProductController struct {
	ProductUsecase productUseCase.Usecase
	UserUsecase    userUsecase.Usecase
	SessionManager *sessionCommon.SessionManager
}

func NewProductController(router *gin.RouterGroup, productUsecase productUseCase.Usecase, userUsecase userUsecase.Usecase, sessionManager *sessionCommon.SessionManager) {
	productController := ProductController{
		ProductUsecase: productUsecase,
		UserUsecase:    userUsecase,
		SessionManager: sessionManager,
	}

	router.Use(httpMiddleware.SessionAuthMiddleware(productController.SessionManager))

	router.POST("/product", productController.AddProduct)
	router.POST("/product/image", productController.AddProductImage)
	router.POST("/cart", productController.AddProductToCart)
	router.POST("/order", productController.CreateOrder)
	router.PUT("/product/:productID", productController.UpdateProduct)
	router.PUT("/order/status", productController.UpdateOrderStatus)
	router.PUT("/cart/:productID/:qty", productController.UpdateProductQtyInCart)
	router.DELETE("/cart/:productID", productController.DeleteProductInCart)
	router.DELETE("/product/:productID", productController.DeleteProduct)
}

func (p *ProductController) CreateOrder(ctx *gin.Context) {
	var req httpCommon.CreateOrder

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(errorCommon.NewInvariantError(err.Error()))
		return
	}

	customer, err := p.UserUsecase.GetCustomerByUserID(ctx, ctx.GetString("user_id"))

	if err != nil {
		ctx.Error(err)
		return
	}

	err = p.ProductUsecase.InsertOrder(ctx, customer.CustomerID, req)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Order created",
	})
}

func (p *ProductController) AddProduct(ctx *gin.Context) {
	var req httpCommon.ProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(errorCommon.NewInvariantError(err.Error()))
		return
	}

	seller, err := p.UserUsecase.GetSellerByUserID(ctx, ctx.GetString("user_id"))

	if err != nil {
		ctx.Error(err)
		return
	}

	err = p.ProductUsecase.InsertProduct(ctx, seller.SellerID, req)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Product added",
	})
}

func (p *ProductController) AddProductImage(ctx *gin.Context) {
	file, err := ctx.FormFile("image")

	if err != nil {
		ctx.Error(errorCommon.NewInvariantError(err.Error()))
		return
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	// The file is received, so let's save it
	if err := ctx.SaveUploadedFile(file, "/public/product_image/"+newFileName); err != nil {
		ctx.Error(errorCommon.NewInvariantError(err.Error()))
		return
	}

	p.ProductUsecase.InsertImage(ctx, newFileName)

	ctx.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Product image added",
	})
}

func (p *ProductController) AddProductToCart(ctx *gin.Context) {
	var req httpCommon.AddCartProduct

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(errorCommon.NewInvariantError(err.Error()))
		return
	}

	customer, err := p.UserUsecase.GetCustomerByUserID(ctx, ctx.GetString("user_id"))

	if err != nil {
		ctx.Error(err)
		return
	}

	err = p.ProductUsecase.InsertCartProduct(ctx, customer.CustomerID, req)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Product added to cart",
	})
}

func (p *ProductController) UpdateProduct(ctx *gin.Context) {
	var req httpCommon.ProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(errorCommon.NewInvariantError(err.Error()))
		return
	}

	productID := ctx.Param("productID")

	seller, err := p.UserUsecase.GetSellerByUserID(ctx, ctx.GetString("user_id"))

	if err != nil {
		ctx.Error(err)
		return
	}

	product, err := p.ProductUsecase.GetProductByID(ctx, productID)

	if err != nil {
		ctx.Error(err)
		return
	}

	if product.SellerID != seller.SellerID {
		ctx.Error(errorCommon.NewForbiddenError("You are not allowed to update this product"))
		return
	}

	err = p.ProductUsecase.UpdateProductByID(ctx, productID, req)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Product updated",
	})
}

func (p *ProductController) UpdateProductQtyInCart(ctx *gin.Context) {
	cartProductID := ctx.Param("productID")
	qty := ctx.Param("qty")

	cartProduct, err := p.ProductUsecase.GetCartProductByID(ctx, cartProductID)

	if err != nil {
		ctx.Error(err)
		return
	}

	customer, err := p.UserUsecase.GetCustomerByUserID(ctx, ctx.GetString("user_id"))

	if err != nil {
		ctx.Error(err)
		return
	}

	if cartProduct.CustomerID != customer.CustomerID {
		ctx.Error(errorCommon.NewForbiddenError("You are not allowed to update this product"))
		return
	}

	err = p.ProductUsecase.UpdateCartProductQtyByID(ctx, cartProductID, qty)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Product quantity updated",
	})
}

func (p *ProductController) UpdateOrderStatus(ctx *gin.Context) {
	var req httpCommon.UpdateOrderStatus

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(errorCommon.NewInvariantError(err.Error()))
		return
	}

	order, err := p.ProductUsecase.GetOrderByID(ctx, req.OrderID)

	if err != nil {
		ctx.Error(err)
		return
	}

	seller, err := p.UserUsecase.GetSellerByUserID(ctx, ctx.GetString("user_id"))

	if err != nil {
		ctx.Error(err)
		return
	}

	if order.SellerID != seller.SellerID {
		ctx.Error(errorCommon.NewForbiddenError("You are not allowed to update this order"))
		return
	}

	err = p.ProductUsecase.UpdateOrderStatusByID(ctx, req)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Order status updated",
	})
}

func (p *ProductController) DeleteProductInCart(ctx *gin.Context) {
	productID := ctx.Param("productID")

	cartProduct, err := p.ProductUsecase.GetCartProductByID(ctx, productID)

	if err != nil {
		ctx.Error(err)
		return
	}

	customer, err := p.UserUsecase.GetCustomerByUserID(ctx, ctx.GetString("user_id"))

	if err != nil {
		ctx.Error(err)
		return
	}

	if cartProduct.CustomerID != customer.CustomerID {
		ctx.Error(errorCommon.NewForbiddenError("You are not allowed to delete this product"))
		return
	}

	err = p.ProductUsecase.DeleteCartProductByID(ctx, productID)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Product deleted",
	})
}

func (p *ProductController) DeleteProduct(ctx *gin.Context) {
	productID := ctx.Param("productID")

	seller, err := p.UserUsecase.GetSellerByUserID(ctx, ctx.GetString("user_id"))

	if err != nil {
		ctx.Error(err)
		return
	}

	product, err := p.ProductUsecase.GetProductByID(ctx, productID)

	if err != nil {
		ctx.Error(err)
		return
	}

	if product.SellerID != seller.SellerID {
		ctx.Error(errorCommon.NewForbiddenError("You are not allowed to delete this product"))
		return
	}

	err = p.ProductUsecase.DeleteProductByID(ctx, productID)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Product deleted",
	})
}
