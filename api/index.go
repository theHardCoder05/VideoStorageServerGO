package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Welcome to Storage Server Service!!!")
}
