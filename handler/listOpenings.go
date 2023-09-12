package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicholasboari/gopportunities/schemas"
)

// @BasePath /api/v1
func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error to list openings")
		return
	}

	sendSuccess(ctx, http.StatusOK, "list-openings", openings)
}
