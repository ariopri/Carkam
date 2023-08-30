package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-66/config"
	"github.com/rg-km/final-project-engineering-66/repository"
	"github.com/rg-km/final-project-engineering-66/utils"
	"net/http"
	"strings"
)

func DeserializeUser(repository repository.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		authorizationHeader := ctx.GetHeader("Authorization")
		fields := strings.Fields(authorizationHeader)
		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		config, _ := config.LoadConfig()
		sub, err := utils.ValidateToken(token, config.TokenSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		userID := sub.(string)
		result, err := repository.FindByID(ctx, nil, userID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		ctx.Set("user", result.UserName)
		ctx.Next()
	}
}
