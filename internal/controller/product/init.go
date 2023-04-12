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

	router.GET("/product", productController.GetRecommendedProduct)
	router.GET("/product/search/:search", productController.GetSearchedProduct)
	router.GET("/product/seller/:sellerID", productController.GetAllProductBySellerID)
	router.GET("/product/:productID", productController.GetProductByID)
	router.GET("/order/:orderID", productController.GetOrderByID)
	router.GET("/order/customer", productController.GetCustomerOrder)
	router.GET("/order/seller", productController.GetSellerOrder)
	router.GET("/order/product/:orderID", productController.GetOrderProductByOrderID)
	router.POST("/order", productController.AddOrder)
	router.PUT("/order/status", productController.ChangeOrderStatus)
	router.GET("/cart", productController.GetCustomerCart)
	router.POST("/cart", productController.AddProductToCart)
	router.PUT("/cart/:productID/:qty", productController.UpdateProductQtyInCart)
	router.DELETE("/cart/:productID", productController.DeleteProductInCart)
	router.POST("/product", productController.AddProduct)
	router.PUT("/product", productController.UpdateProduct)
	router.DELETE("/product/:productID", productController.DeleteProduct)
}

func (p *ProductController) GetRecommendedProduct(ctx *gin.Context) {

}

func (p *ProductController) GetSearchedProduct(ctx *gin.Context) {

}

func (p *ProductController) GetAllProductBySellerID(ctx *gin.Context) {

}

func (p *ProductController) GetProductByID(ctx *gin.Context) {

}

func (p *ProductController) GetOrderByID(ctx *gin.Context) {

}

func (p *ProductController) GetCustomerOrder(ctx *gin.Context) {

}

func (p *ProductController) GetSellerOrder(ctx *gin.Context) {

}

func (p *ProductController) GetOrderProductByOrderID(ctx *gin.Context) {

}

func (p *ProductController) GetCustomerCart(ctx *gin.Context) {

}

func (p *ProductController) AddProductToCart(ctx *gin.Context) {

}

func (p *ProductController) UpdateProductQtyInCart(ctx *gin.Context) {
}

func (p *ProductController) DeleteProductInCart(ctx *gin.Context) {

}

func (p *ProductController) AddOrder(ctx *gin.Context) {

}

func (p *ProductController) AddProduct(ctx *gin.Context) {

}

func (p *ProductController) UpdateProduct(ctx *gin.Context) {

}

func (p *ProductController) DeleteProduct(ctx *gin.Context) {

}

func (p *ProductController) ChangeOrderStatus(ctx *gin.Context) {

}
