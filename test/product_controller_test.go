package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	dbCommon "github.com/aziemp66/byte-bargain/common/db"
	httpCommon "github.com/aziemp66/byte-bargain/common/http"
	sessionCommon "github.com/aziemp66/byte-bargain/common/session"
	productController "github.com/aziemp66/byte-bargain/internal/controller/product"
	productRepository "github.com/aziemp66/byte-bargain/internal/repository/product"
	productUsecase "github.com/aziemp66/byte-bargain/internal/usecase/product"
)

func getDummyProductController() *gin.Engine {
	router := gin.Default()

	gin.SetMode(gin.TestMode)

	productGroup := router.Group("/api/product")

	testDb := dbCommon.NewDB("root:azie122333@tcp(localhost:3306)/test_byte_bargain?charset=utf8mb4&parseTime=True&loc=Local")
	testSession := sessionCommon.NewSessionManager([]byte("secret"))
	ProductRepository := productRepository.NewProductRepositoryImplementation()
	ProductUsecase := productUsecase.NewProductUsecaseImplementation(ProductRepository, testDb, testSession)
	productController.NewProductController(productGroup, ProductUsecase)

	return router
}

func TestProductController(t *testing.T) {
	router := getDummyProductController()

	t.Run("Test Get All Products", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/api/product", nil)

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
	})

	t.Run("Test Get Product By Search", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/api/product/search/yoursearch", nil)

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
	})

	t.Run("Test Get Product By ID", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/api/product/1", nil)

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}

	})

	t.Run("Test Get Product By SellerID", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/api/product/seller/1", nil)

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
	})

	t.Run("Test Get Order By ID", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/api/order/1", nil)

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
	})

	t.Run("Test Get Customer Order", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/api/order/customer", nil)

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
	})

	t.Run("Test Get Seller Order", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/api/order/seller", nil)

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
	})

	t.Run("Test Get Order Product By OrderID", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/api/order/product/1", nil)

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}

	})

	t.Run("Test Get Customer Cart", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/api/cart", nil)

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}

	})

	t.Run("Test Add Product To Cart", func(t *testing.T) {
		requestBody := httpCommon.AddProduct{
			Name:        "Test Product",
			Price:       100000.500,
			Stock:       10,
			Category:    "Test Category",
			Description: "Test Description",
			Weight:      10,
		}

		requestBodyBytes, _ := json.Marshal(requestBody)

		request, _ := http.NewRequest("POST", "/api/cart", bytes.NewBuffer(requestBodyBytes))

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}

	})

	t.Run("Test Update Product Quantity From Cart", func(t *testing.T) {
		request, _ := http.NewRequest("PUT", "/api/cart/12/5", nil)

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
	})

	t.Run("Test Add Product", func(t *testing.T) {
		requestBody := httpCommon.Product{
			SellerID:    "1",
			Name:        "Test Product",
			Price:       100000.500,
			Stock:       10,
			Category:    "Test Category",
			Description: "Test Description",
			Weight:      10,
		}

		requestBodyBytes, _ := json.Marshal(requestBody)

		request, _ := http.NewRequest("POST", "/api/product", bytes.NewBuffer(requestBodyBytes))

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
	})

	t.Run("Test Update Product", func(t *testing.T) {
		requestBody := httpCommon.Product{
			ID:          "1",
			SellerID:    "1",
			Name:        "Test Product",
			Price:       100000.500,
			Stock:       10,
			Category:    "Test Category",
			Description: "Test Description",
			Weight:      10,
		}

		requestBodyBytes, _ := json.Marshal(requestBody)

		request, _ := http.NewRequest("PUT", "/api/product", bytes.NewBuffer(requestBodyBytes))

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
	})

	t.Run("Test Delete Product", func(t *testing.T) {
		request, _ := http.NewRequest("DELETE", "/api/product/1", nil)

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
	})

	t.Run("Test Add Order", func(t *testing.T) {
		requestBody := httpCommon.CreateOrder{
			SellerID: "1",
			Items: []httpCommon.ProductItem{
				{ProductID: "1", Qty: 1},
			},
		}

		requestBodyBytes, _ := json.Marshal(requestBody)

		request, _ := http.NewRequest("POST", "/api/order", bytes.NewBuffer(requestBodyBytes))

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
	})

	t.Run("Test Change Order Status", func(t *testing.T) {
		requestBody := httpCommon.UpdateOrderStatus{
			OrderID: "1",
			Status:  "Success",
		}

		requestBodyBytes, _ := json.Marshal(requestBody)

		request, _ := http.NewRequest("PUT", "/api/order/status", bytes.NewBuffer(requestBodyBytes))

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
	})

}
