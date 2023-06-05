package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/auth"
	authpb "github.com/Portfolio-Advanced-software/API-Gateway/pkg/auth/pb"
	mockauthpb "github.com/Portfolio-Advanced-software/API-Gateway/pkg/auth/pb/mock"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/authz"
	authzpb "github.com/Portfolio-Advanced-software/API-Gateway/pkg/authz/pb"
	mockauthzpb "github.com/Portfolio-Advanced-software/API-Gateway/pkg/authz/pb/mock"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/config"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/history"
	historypb "github.com/Portfolio-Advanced-software/API-Gateway/pkg/history/pb"
	mockhistorypb "github.com/Portfolio-Advanced-software/API-Gateway/pkg/history/pb/mock"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/movie"
	moviepb "github.com/Portfolio-Advanced-software/API-Gateway/pkg/movie/pb"
	mockmoviepb "github.com/Portfolio-Advanced-software/API-Gateway/pkg/movie/pb/mock"
	"github.com/gin-gonic/gin"
)

func TestHistoryAPI(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthService := mockauthpb.NewMockAuthServiceClient(ctrl)
	mockAuthzService := mockauthzpb.NewMockAuthzServiceClient(ctrl)
	mockHistoryService := mockhistorypb.NewMockHistoryServiceClient(ctrl)
	historyService := history.ServiceClient{Client: mockHistoryService}
	mockMovieService := mockmoviepb.NewMockMovieServiceClient(ctrl)
	movieService := movie.ServiceClient{Client: mockMovieService}

	cfg := &config.Config{
		// Set your configuration values here
	}

	authSvc := auth.ServiceClient{Client: mockAuthService}
	authzSvc := authz.ServiceClient{Client: mockAuthzService}
	//historySvc := history.ServiceClient{Client: mockhistorypb.NewMockHistoryServiceClient(ctrl)}

	router := gin.New()

	auth.RegisterRoutes(router, cfg)
	authz.RegisterRoutes(router, cfg)
	history.RegisterRoutes(router, cfg, &authSvc, &authzSvc, &historyService)
	movie.RegisterRoutes(router, cfg, &authSvc, &authzSvc, &movieService)

	t.Run("Test ReadMovie with valid JWT token", func(t *testing.T) {
		// Mock the authentication service to return a valid token
		mockAuthService.EXPECT().Validate(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, req *authpb.ValidateRequest, opts ...grpc.CallOption) (*authpb.ValidateResponse, error) {
			return &authpb.ValidateResponse{
				Status: http.StatusOK, // Mocking a non-admin role
				Error:  "",
				UserId: "1",
			}, nil
		})

		// Mock the authentication service to return a valid token
		mockMovieService.EXPECT().ReadMovie(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, req *moviepb.ReadMovieReq, opts ...grpc.CallOption) (*moviepb.ReadMovieRes, error) {
			return &moviepb.ReadMovieRes{
				Movie: &moviepb.Movie{
					Title: "avengers",
				},
			}, nil
		})

		// Generate a new ObjectID
		objectID := primitive.NewObjectID()
		req := httptest.NewRequest(http.MethodGet, "/movie/"+objectID.Hex(), nil)
		//req.Header.Set("Content-Type", "application/json")
		req.Header.Set("authorization", "Bearer valid-token")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		// Perform assertions on the response body or other details
		// ...
	})

	t.Run("Test ReadHistory without JWT token", func(t *testing.T) {
		// Generate a new ObjectID
		objectID := primitive.NewObjectID()
		req := httptest.NewRequest(http.MethodGet, "/movie/"+objectID.Hex(), nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)

		// Perform assertions on the response body or other details
		// ...
	})

	t.Run("Test ReadHistory with admin role", func(t *testing.T) {
		// Mock the authentication service to return a valid token
		mockAuthService.EXPECT().Validate(gomock.Any(), gomock.Any()).Times(1).DoAndReturn(func(ctx context.Context, req *authpb.ValidateRequest, opts ...grpc.CallOption) (*authpb.ValidateResponse, error) {
			return &authpb.ValidateResponse{
				Status: http.StatusOK,
				Error:  "",
				UserId: "1",
			}, nil
		})

		// Mock the authentication service to return a valid token
		mockAuthzService.EXPECT().VerifyRole(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, req *authzpb.VerifyRoleRequest, opts ...grpc.CallOption) (*authzpb.VerifyRoleResponse, error) {
			return &authzpb.VerifyRoleResponse{
				IsAuthorized: true,
			}, nil
		})

		// Mock the authentication service to return a valid token
		mockHistoryService.EXPECT().ReadHistory(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, req *historypb.ReadHistoryReq, opts ...grpc.CallOption) (*historypb.ReadHistoryRes, error) {
			return &historypb.ReadHistoryRes{
				History: &historypb.History{
					Userid:  "23535",
					Movieid: "214234",
				},
			}, nil
		})

		// Generate a new ObjectID
		objectID := primitive.NewObjectID()
		req := httptest.NewRequest(http.MethodGet, "/history/"+objectID.Hex(), nil)
		//req.Header.Set("Content-Type", "application/json")
		req.Header.Set("authorization", "Bearer valid-token")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("Test ReadHistory without admin role", func(t *testing.T) {
		// Mock the authentication service to return a valid token
		mockAuthService.EXPECT().Validate(gomock.Any(), gomock.Any()).Times(1).DoAndReturn(func(ctx context.Context, req *authpb.ValidateRequest, opts ...grpc.CallOption) (*authpb.ValidateResponse, error) {
			return &authpb.ValidateResponse{
				Status: http.StatusOK,
				Error:  "",
				UserId: "1",
			}, nil
		})

		// Mock the authentication service to return a valid token
		mockAuthzService.EXPECT().VerifyRole(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, req *authzpb.VerifyRoleRequest, opts ...grpc.CallOption) (*authzpb.VerifyRoleResponse, error) {
			return &authzpb.VerifyRoleResponse{
				IsAuthorized: false,
			}, nil
		})

		// Generate a new ObjectID
		objectID := primitive.NewObjectID()
		req := httptest.NewRequest(http.MethodGet, "/history/"+objectID.Hex(), nil)
		//req.Header.Set("Content-Type", "application/json")
		req.Header.Set("authorization", "Bearer valid-token")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusForbidden, w.Code)

	})
}
