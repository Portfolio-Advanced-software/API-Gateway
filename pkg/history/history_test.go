package history

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateHistoryWithoutToken(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Register the routes
	RegisterRoutes(router, nil, nil) // Pass nil as the authSvc since we're testing without token

	// Create an HTTP request to the CreateHistory endpoint
	req, err := http.NewRequest(http.MethodPost, "/history/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create an HTTP test recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the request using the router
	router.ServeHTTP(recorder, req)

	// Assert that the response code is 401 Unauthorized
	if recorder.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d but got %d", http.StatusUnauthorized, recorder.Code)
	}
}

func TestCreateHistoryWithToken(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Register the routes
	RegisterRoutes(router, nil, nil) // Pass nil as the authSvc since we're testing without token

	// Create an HTTP request to the CreateHistory endpoint
	req, err := http.NewRequest(http.MethodPost, "/history/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set the JWT token in the request headers
	req.Header.Set("Authorization", "Bearer your-jwt-token")

	// Create an HTTP test recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the request using the router
	router.ServeHTTP(recorder, req)

	// Assert that the response code is not 401 Unauthorized
	if recorder.Code == http.StatusUnauthorized {
		t.Errorf("Expected a response other than %d", http.StatusUnauthorized)
	}
}
