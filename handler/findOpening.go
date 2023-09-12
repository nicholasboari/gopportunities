package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicholasboari/gopportunities/schemas"
)

func FindOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "string").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	sendSuccess(ctx, http.StatusOK, "show-opening", opening)
}
