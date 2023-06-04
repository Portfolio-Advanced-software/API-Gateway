package auth_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"net"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/auth/pb"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/auth/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

type AuthServiceServerMock struct {
	pb.UnimplementedAuthServiceServer
}

type UserMock struct {
	Id       primitive.ObjectID
	Email    string
	Password string
}

func (m *AuthServiceServerMock) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if req.Email == "existing@example.com" {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "E-Mail already exists",
		}, nil
	}

	// Simulate successful registration
	insertedUserID := primitive.NewObjectID()
	fmt.Println("Created user with id: " + insertedUserID.Hex())

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (m *AuthServiceServerMock) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req.Email == "nonexisting@example.com" {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	// Simulate successful login
	user := UserMock{
		Email:    req.Email,
		Password: "hashed_password",
	}
	fmt.Println(user.Email + " logged in.")

	token := "generated_token"

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func TestRegister(t *testing.T) {
	// Create a new test gRPC server
	testServer := grpc.NewServer()

	// Create an instance of the AuthServiceServerMock
	authServer := &AuthServiceServerMock{}

	// Register the AuthServiceServerMock as the gRPC server implementation
	pb.RegisterAuthServiceServer(testServer, authServer)

	// Start the test server in a separate goroutine
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("failed to start test server: %v", err)
	}
	go func() {
		if err := testServer.Serve(lis); err != nil {
			t.Fatalf("failed to start test server: %v", err)
		}
	}()

	// Create a new gRPC connection to the test server
	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to connect to test server: %v", err)
	}
	defer conn.Close()

	// Create an instance of the AuthServiceClient using the test server connection
	authClient := pb.NewAuthServiceClient(conn)

	// Create a new Gin router
	router := gin.Default()

	// Register the register route
	router.POST("/register", func(c *gin.Context) {
		// Call the Register handler
		routes.Register(c, authClient)
	})

	t.Run("successful registration", func(t *testing.T) {
		// Create a test request body with unique registration details
		body := routes.RegisterRequestBody{
			Email:    "test@example.com",
			Password: "password",
		}
		jsonBody, _ := json.Marshal(body)

		// Create a test HTTP request with the JSON body
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Create a test HTTP response recorder
		res := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(res, req)

		// Assert that the response has the expected status code
		assert.Equal(t, http.StatusCreated, res.Code)
	})

	t.Run("duplicate registration", func(t *testing.T) {
		// Create a test request body with duplicate registration details
		body := routes.RegisterRequestBody{
			Email:    "existing@example.com",
			Password: "password",
		}
		jsonBody, _ := json.Marshal(body)

		// Create a test HTTP request with the JSON body
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Create a test HTTP response recorder
		res := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(res, req)

		// Assert that the response has the expected status code
		assert.Equal(t, http.StatusConflict, res.Code)
	})
}

func TestLogin(t *testing.T) {
	// Create a new test gRPC server
	testServer := grpc.NewServer()

	// Create an instance of the AuthServiceServerMock
	authServer := &AuthServiceServerMock{}

	// Register the AuthServiceServerMock as the gRPC server implementation
	pb.RegisterAuthServiceServer(testServer, authServer)

	// Start the test server in a separate goroutine
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("failed to start test server: %v", err)
	}
	go func() {
		if err := testServer.Serve(lis); err != nil {
			t.Fatalf("failed to start test server: %v", err)
		}
	}()

	// Create a new gRPC connection to the test server
	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to connect to test server: %v", err)
	}
	defer conn.Close()

	// Create an instance of the AuthServiceClient using the test server connection
	authClient := pb.NewAuthServiceClient(conn)

	// Create a new Gin router
	router := gin.Default()

	// Register the login route
	router.POST("/login", func(c *gin.Context) {
		// Call the Login handler
		routes.Login(c, authClient)
	})

	t.Run("successful login", func(t *testing.T) {
		// Create a test request body with valid login details
		body := routes.LoginRequestBody{
			Email:    "user@example.com",
			Password: "password",
		}
		jsonBody, _ := json.Marshal(body)

		// Create a test HTTP request with the JSON body
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Create a test HTTP response recorder
		res := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(res, req)

		// Assert that the response has the expected status code
		assert.Equal(t, http.StatusCreated, res.Code)
	})

	t.Run("invalid login", func(t *testing.T) {
		// Create a test request body with invalid login details
		body := routes.LoginRequestBody{
			Email:    "nonexisting@example.com",
			Password: "password",
		}
		jsonBody, _ := json.Marshal(body)

		// Create a test HTTP request with the JSON body
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Create a test HTTP response recorder
		res := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(res, req)

		// Assert that the response has the expected status code
		assert.Equal(t, http.StatusNotFound, res.Code)
	})
}
