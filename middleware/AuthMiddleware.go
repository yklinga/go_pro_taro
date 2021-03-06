package middleware

import (
	"github.com/gin-gonic/gin"
	"go_pro_taro/model"
	"go_pro_taro/utils"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")

		if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "token失效"})
			ctx.Abort()
			return
		}

		tokenStr = tokenStr[7:]

		token, claims, err := utils.ParseToken(tokenStr)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "token失效"})
			ctx.Abort()
			return
		}

		// 通过验证后 获取 claim 中的 userId

		userId := claims.UserId
		DB :=utils.GetDB()
		var user model.User
		DB.First(&user, userId)

		//	用户 不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "user_id 获取失败，用户不存在"})
			ctx.Abort()
			return
		}

		//	用户存在 将user写入上下文

		ctx.Set("user", user)

		ctx.Next()
	}
}
