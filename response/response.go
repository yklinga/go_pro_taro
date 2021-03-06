package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResCommon(ctx * gin.Context, httpStatus int, code int, data gin.H, msg string)  {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(ctx * gin.Context, data gin.H, msg string)  {
	if msg == "" {
		msg = "success"
	}
	ResCommon(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx * gin.Context, data gin.H, msg string)  {
	ResCommon(ctx, http.StatusBadRequest, 400, data, msg)
}
