package product

import (
	"github.com/gin-gonic/gin"

	errorCommon "github.com/aziemp66/byte-bargain/common/error"
	httpCommon "github.com/aziemp66/byte-bargain/common/http"

	productUseCase "github.com/aziemp66/byte-bargain/internal/usecase/product"
)

type ProductController struct {
	ProductUsecase productUseCase.Usecase
}

func NewProductController(router *gin.RouterGroup, productUsecase productUseCase.Usecase) {
	productController := ProductController{
		ProductUsecase: productUsecase,
	}

	router.POST("/product", productController.AddProduct)
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

	err := p.ProductUsecase.InsertOrder(ctx, req)

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

	err := p.ProductUsecase.InsertProduct(ctx, req)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Product added",
	})
}
func (p *ProductController) AddProductToCart(ctx *gin.Context) {
	var req httpCommon.AddCartProduct

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(errorCommon.NewInvariantError(err.Error()))
		return
	}

	err := p.ProductUsecase.InsertCartProduct(ctx, req)

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

	err := p.ProductUsecase.UpdateProductByID(ctx, productID, req)

	if err != nil {
		ctx.Error(err)
		return
	}
}

func (p *ProductController) UpdateProductQtyInCart(ctx *gin.Context) {
	productID := ctx.Param("productID")
	qty := ctx.Param("qty")

	err := p.ProductUsecase.UpdateCartProductQtyByID(ctx, productID, qty)

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

	err := p.ProductUsecase.UpdateOrderStatusByID(ctx, req)

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

	err := p.ProductUsecase.DeleteCartProductByID(ctx, productID)

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

	err := p.ProductUsecase.DeleteProductByID(ctx, productID)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Product deleted",
	})
}
