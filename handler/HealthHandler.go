package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})

	sendSuccess(ctx, "health-check", nil)
}
