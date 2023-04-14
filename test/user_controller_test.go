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
	passwordCommon "github.com/aziemp66/byte-bargain/common/password"
	sessionCommon "github.com/aziemp66/byte-bargain/common/session"

	userController "github.com/aziemp66/byte-bargain/internal/controller/user"
	userRepository "github.com/aziemp66/byte-bargain/internal/repository/user"
	userUsecase "github.com/aziemp66/byte-bargain/internal/usecase/user"
)

func getDummyUserController() *gin.Engine {
	gin.SetMode(gin.TestMode)

	router := gin.Default()

	userGroup := router.Group("/api/user")

	testDb := dbCommon.NewDB("root:azie122333@tcp(localhost:3306)/test_byte_bargain?charset=utf8mb4&parseTime=True&loc=Local")
	testSession := sessionCommon.NewSessionManager([]byte("secret"))
	passwordManager := passwordCommon.NewPasswordHashManager()
	UserRepository := userRepository.NewUserRepositoryImplementation()
	UserUsecase := userUsecase.NewUserUsecaseImplementation(UserRepository, testDb, testSession, passwordManager)
	userController.NewUserController(userGroup, UserUsecase)
	return router
}

func TestUserController(t *testing.T) {
	router := getDummyUserController()

	t.Run("Test Login", func(t *testing.T) {
		requestBody := httpCommon.Login{
			Email:    "youremail@gmail.com",
			Password: "yourpassword",
		}
		requestBodyBytes, _ := json.Marshal(requestBody)
		request, _ := http.NewRequest("POST", "/api/user/login", bytes.NewBuffer(requestBodyBytes))

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
	})

	t.Run("Test Register Customer", func(t *testing.T) {
		requestBody := httpCommon.RegisterCustomer{
			Email:       "youremail@gmail.com",
			Password:    "yourpassword",
			Name:        "yourname",
			Address:     "youraddress",
			BirthDate:   "yourbirthdate",
			PhoneNumber: "yourphonenumber",
			Gender:      "yourgender",
		}

		requestBodyBytes, _ := json.Marshal(requestBody)
		request, _ := http.NewRequest("POST", "/api/user/register/customer", bytes.NewBuffer(requestBodyBytes))

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
	})

	t.Run("Test Register Seller", func(t *testing.T) {
		requestBody := httpCommon.RegisterSeller{
			Email:          "myemail@gmail.com",
			Password:       "mypassword",
			Name:           "myname",
			Address:        "myaddress",
			BirthDate:      "mybirthdate",
			PhoneNumber:    "myphonenumber",
			Gender:         "mygender",
			IdentityNumber: "myidentitynumber",
			BankName:       "mybankname",
			DebitNumber:    "mydebitnumber",
		}

		requestBodyBytes, _ := json.Marshal(requestBody)

		request, _ := http.NewRequest("POST", "/api/user/register/seller", bytes.NewBuffer(requestBodyBytes))

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)
	})

	t.Run("Test Get Customer By ID", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/api/user/customer/1", nil)

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
	})

	t.Run("Test Get Seller By ID", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/api/user/seller/1", nil)

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
	})

	t.Run("Test Forgot Password", func(t *testing.T) {
		requestBody := httpCommon.ForgotPassword{
			Email: "youremail@gmail.com",
		}

		requestBodyBytes, _ := json.Marshal(requestBody)

		request, _ := http.NewRequest("POST", "/api/user/forgot-password", bytes.NewBuffer(requestBodyBytes))

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
	})

	t.Run("Test Reset Password", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/api/user/reset-password/12/yourtoken", nil)

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
	})

	t.Run("Test Change Password", func(t *testing.T) {
		requestBody := httpCommon.ChangePassword{
			OldPassword: "youroldpassword",
			NewPassword: "yournewpassword",
		}

		requestBodyBytes, _ := json.Marshal(requestBody)

		request, _ := http.NewRequest("POST", "/api/user/change-password", bytes.NewBuffer(requestBodyBytes))

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
	})

	t.Run("Test Update Customer", func(t *testing.T) {
		requestBody := httpCommon.UpdateCustomer{
			Name:        "yourname",
			Address:     "youraddress",
			BirthDate:   "yourbirthdate",
			PhoneNumber: "yourphonenumber",
			Gender:      "yourgender",
		}

		requestBodyBytes, _ := json.Marshal(requestBody)

		request, _ := http.NewRequest("POST", "/api/user/customer", bytes.NewBuffer(requestBodyBytes))

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
	})

	t.Run("Test Update Seller", func(t *testing.T) {
		requestBody := httpCommon.UpdateSeller{
			Name:           "yourname",
			Address:        "youraddress",
			BirthDate:      "yourbirthdate",
			PhoneNumber:    "yourphonenumber",
			Gender:         "yourgender",
			IdentityNumber: "youridentitynumber",
			BankName:       "yourbankname",
			DebitNumber:    "yourdebitnumber",
		}

		requestBodyBytes, _ := json.Marshal(requestBody)

		request, _ := http.NewRequest("POST", "/api/user/seller", bytes.NewBuffer(requestBodyBytes))

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
	})
}
