package middlewares

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/fahmyabdul/self-growth/fetch-app/internal/models"
	"github.com/fahmyabdul/self-growth/fetch-app/internal/models/endpoints_permission"
	"github.com/fahmyabdul/self-growth/fetch-app/internal/requests"
	"github.com/gin-gonic/gin"
)

func GetHeaderToken(c *gin.Context) string {
	if _, ok := c.Request.Header["Authorization"]; !ok {
		return ""
	}

	authHeader := c.Request.Header["Authorization"][0]

	if !strings.Contains(authHeader, "Bearer ") {
		return ""
	}

	return strings.Split(authHeader, " ")[1]
}

func JwtAuth(basePath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken := GetHeaderToken(c)

		// If jwt token is empty then return unauthorized
		if jwtToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ResponseRestApi{
				Status: "Unauthorized",
				Code:   http.StatusUnauthorized,
				Data:   "",
			})
			return
		}

		// Validate jwt token to auth-app
		var authApp = requests.AuthApp{}
		responseData, err := authApp.Validate(jwtToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ResponseRestApi{
				Status: "Unauthorized",
				Code:   http.StatusUnauthorized,
				Data:   "",
			})
			return
		}

		// If jwt token is invalid/expired then return unauthorized
		if responseData.Message != "Valid JWT" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ResponseRestApi{
				Status: "Unauthorized",
				Code:   http.StatusUnauthorized,
				Data:   "",
			})
			return
		}

		jsonPayload, err := json.Marshal(responseData)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ResponseRestApi{
				Status: "Unauthorized",
				Code:   http.StatusUnauthorized,
				Data:   "",
			})
			return
		}

		// Get Permission List
		var modelPermission endpoints_permission.EndpointsPermission
		listPermission, err := modelPermission.GetListPermission()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ResponseRestApi{
				Status: "Unauthorized",
				Code:   http.StatusUnauthorized,
				Data:   "",
			})
			return
		}

		path := strings.ReplaceAll(c.FullPath(), basePath, "")
		permission := listPermission[path]
		if permission != "*" && responseData.Data.Role != permission {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ResponseRestApi{
				Status: "Unauthorized",
				Code:   http.StatusUnauthorized,
				Data:   "",
			})
		}

		c.Writer.Header().Set("x-data-user", string(jsonPayload))
		c.Next()
	}
}
