package gu

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type PageData struct {
	Total    int64 `json:"total"`
	PageID   int64 `json:"page_id"`
	PageSize int64 `json:"page_size"`
	Result   any   `json:"result"`
}

func Ok(ctx *gin.Context, data any) {
	ResponseJson(ctx, http.StatusOK, "ok", data)
}

func Failed(ctx *gin.Context, message string) {
	ResponseJson(ctx, http.StatusInternalServerError, message, nil)
}

func FailedWithCode(ctx *gin.Context, code int, message string) {
	ResponseJson(ctx, code, message, nil)
}

func ResponseJson(ctx *gin.Context, code int, message string, data any) {
	ctx.JSON(code, response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
