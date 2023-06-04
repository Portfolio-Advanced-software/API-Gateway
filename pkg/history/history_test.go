package history_test

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"google.golang.org/grpc"

// 	authpb "github.com/Portfolio-Advanced-software/API-Gateway/pkg/auth/pb"
// 	amock "github.com/Portfolio-Advanced-software/API-Gateway/pkg/auth/pb/mock"
// 	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/config"
// 	historypb "github.com/Portfolio-Advanced-software/API-Gateway/pkg/history/pb"
// 	hmock "github.com/Portfolio-Advanced-software/API-Gateway/pkg/history/pb/mock"
// )

// func TestGetHistory(t *testing.T) {
// 	// Create a new Gin router
// 	router := gin.Default()

// 	// Create a new config
// 	c := config.Config{
// 		Port: "8080", // Set the desired port
// 	}

// 	// Create an instance of the AuthServiceServerMock
// 	authServer := amock.MockAuthServiceServer{}
// 	// Create an instance of the HistoryServiceServerMock
// 	historyServer := hmock.MockHistoryServiceServer{}

// 	// Register the AuthServiceServerMock as the gRPC server implementation
// 	authpb.RegisterAuthServiceServer(historyServer, authServer)

// 	// Register the HistoryServiceServerMock as the gRPC server implementation
// 	historypb.RegisterHistoryServiceServer(historyServer, historyServer)

// 	// Start the test gRPC server
// 	testServer := grpc.NewServer()
// 	defer testServer.Stop()

// 	// Register the mock gRPC servers with the test gRPC server
// 	authpb.RegisterAuthServiceServer(testServer, authServer)
// 	historypb.RegisterHistoryServiceServer(testServer, historyServer)

// 	// Start the test server in a separate goroutine
// 	go func() {
// 		if err := router.Run(c.Port); err != nil {
// 			t.Fatalf("failed to start test server: %v", err)
// 		}
// 	}()

// 	// Create a new gRPC connection to the test server
// 	conn, err := grpc.Dial("localhost"+c.Port, grpc.WithInsecure())
// 	if err != nil {
// 		t.Fatalf("failed to connect to test server: %v", err)
// 	}
// 	defer conn.Close()

// 	// Create an instance of the AuthServiceClient using the test server connection
// 	authClient := authpb.NewAuthServiceClient(conn)

// 	// Create a new HTTP request to the test server
// 	req, err := http.NewRequest("GET", "/history", nil)
// 	if err != nil {
// 		t.Fatalf("failed to create HTTP request: %v", err)
// 	}

// 	// Set any required headers or parameters for the request

// 	// Perform the HTTP request
// 	resp := httptest.NewRecorder()
// 	router.ServeHTTP(resp, req)

// 	// Assert the response status code and any other expected values

// 	assert.Equal(t, http.StatusOK, resp.Code, "unexpected status code")
// 	// Assert other response properties or content
// }
