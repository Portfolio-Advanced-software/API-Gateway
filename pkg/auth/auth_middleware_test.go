package auth_test

// import (
// 	"context"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/auth"
// 	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/auth/pb"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type AuthServiceClientMock struct {
// 	auth.ServiceClient
// 	authServer *AuthServiceServerMock
// }

// type CustomServiceClient struct {
// 	auth.ServiceClient
// }

// // func (m *AuthServiceServerMock) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
// // 	if req.Token == "invalid_token" {
// // 		return &pb.ValidateResponse{
// // 			Status: http.StatusBadRequest,
// // 			Error:  "Invalid token",
// // 		}, nil
// // 	}

// // 	// Simulate successful validation
// // 	user := UserMock{
// // 		Id:    primitive.NewObjectID(),
// // 		Email: "user@example.com",
// // 	}

// // 	return &pb.ValidateResponse{
// // 		Status: http.StatusOK,
// // 		UserId: user.Id.Hex(),
// // 	}, nil
// // }

// func (m *AuthServiceClientMock) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
// 	return m.authServer.Validate(ctx, req)
// }

// func (c *CustomServiceClient) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
// 	authClient := &AuthServiceClientMock{authServer: &AuthServiceServerMock{}}
// 	return authClient.Validate(ctx, req)
// }

// func TestAuthRequired(t *testing.T) {
// 	// Create a new Gin router
// 	router := gin.Default()

// 	// Create an instance of the CustomServiceClient
// 	serviceClient := &CustomServiceClient{}

// 	a := auth.InitAuthMiddleware(serviceClient)

// 	// Register the auth required middleware
// 	router.Use(a.AuthRequired)

// 	// Register a test route that requires authentication
// 	router.GET("/protected", func(c *gin.Context) {
// 		userId, _ := c.Get("userId")
// 		c.JSON(http.StatusOK, gin.H{
// 			"userId": userId,
// 		})
// 	})

// 	t.Run("valid token", func(t *testing.T) {
// 		// Create a test HTTP request with a valid token in the authorization header
// 		req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
// 		req.Header.Set("Authorization", "Bearer valid_token")

// 		// Create a test HTTP response recorder
// 		res := httptest.NewRecorder()

// 		// Perform the request
// 		router.ServeHTTP(res, req)

// 		// Assert that the response has the expected status code
// 		assert.Equal(t, http.StatusOK, res.Code)

// 		// Assert that the response body contains the expected user ID
// 		assert.Equal(t, `{"userId":"123"}`, res.Body.String())
// 	})

// 	t.Run("invalid token", func(t *testing.T) {
// 		// Create a test HTTP request with an invalid token in the authorization header
// 		req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
// 		req.Header.Set("Authorization", "Bearer invalid_token")

// 		// Create a test HTTP response recorder
// 		res := httptest.NewRecorder()

// 		// Perform the request
// 		router.ServeHTTP(res, req)

// 		// Assert that the response has the expected status code
// 		assert.Equal(t, http.StatusUnauthorized, res.Code)
// 	})

// 	t.Run("missing token", func(t *testing.T) {
// 		// Create a test HTTP request without an authorization header
// 		req, _ := http.NewRequest(http.MethodGet, "/protected", nil)

// 		// Create a test HTTP response recorder
// 		res := httptest.NewRecorder()

// 		// Perform the request
// 		router.ServeHTTP(res, req)

// 		// Assert that the response has the expected status code
// 		assert.Equal(t, http.StatusUnauthorized, res.Code)
// 	})
// }
