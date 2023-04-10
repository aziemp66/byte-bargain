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
	UserRepository := userRepository.NewUserRepositoryImplementation()
	UserUsecase := userUsecase.NewUserUsecaseImplementation(UserRepository, testDb, testSession)
	userController.NewUserController(userGroup, UserUsecase)
	return router
}

func TestUserController(t *testing.T) {
	router := getDummyUserController()

	t.Run("Test Login", func(t *testing.T) {
		requestBody := map[string]interface{}{
			"email":    "aziemp66@gmail.com",
			"password": "azie122333",
		}
		requestBodyBytes, _ := json.Marshal(requestBody)
		request, _ := http.NewRequest("POST", "/api/user/login", bytes.NewBuffer(requestBodyBytes))

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if assert.NotEqual(t, http.StatusOK, response.Code) {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
		}
	})

}
