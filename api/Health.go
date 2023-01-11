package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck Response Structure
type HealthResponse struct {
	status string
}

//Health Check endpoint
func (server *Server) HealthCheck(ctx *gin.Context) {
	var response HealthResponse
	response.status = "Healthy"
	ctx.JSON(http.StatusOK, gin.H{"Status": response.status})
}
