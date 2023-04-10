package product

import (
	"github.com/gin-gonic/gin"

	productUseCase "github.com/aziemp66/byte-bargain/internal/usecase/product"
)

type ProductController struct {
	ProductUsecase productUseCase.Usecase
}

func NewProductController(productUsecase productUseCase.Usecase) *ProductController {
	return &ProductController{
		ProductUsecase: productUsecase,
	}
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

func (p *ProductController) GetCustomerOrderByID(ctx *gin.Context) {

}

func (p *ProductController) GetSellerOrderByID(ctx *gin.Context) {

}

func (p *ProductController) GetOrderProductByID(ctx *gin.Context) {

}

func (p *ProductController) GetCustomerCart(ctx *gin.Context) {

}

func (p *ProductController) AddProductToCart(ctx *gin.Context) {

}

func (p *ProductController) RemoveProductFromCart(ctx *gin.Context) {

}

func (p *ProductController) Checkout(ctx *gin.Context) {

}

func (p *ProductController) AddProduct(ctx *gin.Context) {

}

func (p *ProductController) UpdateProduct(ctx *gin.Context) {

}

func (p *ProductController) DeleteProduct(ctx *gin.Context) {

}

func (p *ProductController) ChangeOrderStatus(ctx *gin.Context) {

}
