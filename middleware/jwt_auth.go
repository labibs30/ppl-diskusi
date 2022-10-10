package middleware

import (
	"dataekspor-be/common"
	"dataekspor-be/entity"
	"dataekspor-be/service"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthorizeJWT(jwtService service.JWTService, roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := common.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if !strings.Contains(authHeader, "Bearer ") {
			response := common.BuildErrorResponse("Invalid Token", "Token is invalid", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			response := common.BuildErrorResponse("Invalid Token", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if !token.Valid {
			log.Println(err)
			response := common.BuildErrorResponse("Token is invalid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		isAuthorized := false
		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			response := common.BuildErrorResponse("Unknown request", "cant convert claims to map", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		var rolesToCheck entity.SliceOfString

		roleRequests, ok := claims["role"].([]any)
		if !ok {
			response := common.BuildErrorResponse("Unknown request", "cant convert role to slice", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		for _, r := range roleRequests {
			rolesToCheck = append(rolesToCheck, r.(string))
		}

		for _, roleToCheck := range rolesToCheck {
			for _, role := range roles {
				fmt.Println(roleToCheck, "vs", role)
				if roleToCheck == role {
					isAuthorized = true
					break
				}
			}
		}

		fmt.Println("STATUS Authorization:", isAuthorized)

		if !isAuthorized {
			fmt.Println(claims)
			log.Println(err)
			response := common.BuildErrorResponse("Access denied", "You dont have access", nil)
			c.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}

		// fmt.Println("Claim[user_id]", claims["user_id"])
		// fmt.Println("Claim[issuer]", claims["issuer"])
	}
}
