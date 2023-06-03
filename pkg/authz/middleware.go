package authz

import (
	"context"
	"fmt"
	"net/http"

	authzpb "github.com/Portfolio-Advanced-software/API-Gateway/pkg/authz/pb"
	"github.com/gin-gonic/gin"
)

type AuthzMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthzMiddleware(svc *ServiceClient) AuthzMiddlewareConfig {
	return AuthzMiddlewareConfig{svc}
}

func (c *AuthzMiddlewareConfig) RoleRequired(requiredRole string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, exists := ctx.Get("userId")
		if !exists {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Debug statement 1
		fmt.Println("UserID from context:", userID)

		userIDStr, ok := userID.(string)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Debug statement 2
		fmt.Println("UserID string:", userIDStr)

		res, err := c.svc.Client.VerifyRole(context.Background(), &authzpb.VerifyRoleRequest{
			UserId: userIDStr,
			Role:   requiredRole,
		})
		if err != nil || !res.IsAuthorized {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		// Proceed to the next middleware or route handler
		ctx.Next()
	}
}
