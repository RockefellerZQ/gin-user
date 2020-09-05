package middleware

import (
	"gin-user/common"
	"gin-user/dto"
	"gin-user/model"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" ||  !strings.HasPrefix(tokenString, "Bearer "){
			 dto.Failed(ctx, 401, "权限不足")
			 ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseTime(tokenString)
		if err != nil || !token.Valid {
			dto.Failed(ctx, 401, "权限不足")
			ctx.Abort()
			return
		}
		userId := claims.UserId
		var user model.User
		model.GetDB().First(&user, userId)
		if user.ID == 0 {
			dto.Failed(ctx, 401, "权限不足")
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}