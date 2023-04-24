package product

import (
	"github.com/gin-gonic/gin"

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
	router.POST("/order", productController.AddOrder)
	router.PUT("/product", productController.UpdateProduct)
	router.PUT("/order/status", productController.UpdateOrderStatus)
	router.PUT("/cart/:productID/:qty", productController.UpdateProductQtyInCart)
	router.DELETE("/cart/:productID", productController.DeleteProductInCart)
	router.DELETE("/product/:productID", productController.DeleteProduct)
}

func (p *ProductController) AddOrder(ctx *gin.Context) {

}

func (p *ProductController) AddProduct(ctx *gin.Context) {

}
func (p *ProductController) AddProductToCart(ctx *gin.Context) {

}

func (p *ProductController) UpdateProduct(ctx *gin.Context) {

}

func (p *ProductController) UpdateProductQtyInCart(ctx *gin.Context) {
}

func (p *ProductController) UpdateOrderStatus(ctx *gin.Context) {
}

func (p *ProductController) DeleteProductInCart(ctx *gin.Context) {
}

func (p *ProductController) DeleteProduct(ctx *gin.Context) {
}
