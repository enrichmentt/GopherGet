package api

import (
	"net/http"

	"github.com/enrichmentt/GopherGet/internal/models"
	"github.com/gin-gonic/gin"
)

func Information(c *gin.Context) {
	resp := models.InformationResponse{}
	resp.Data.SourceIdentifier = "GopherGet"
	resp.Data.ServerSupportedVersion = []string{"1.0.0", "1.4.0", "1.7.0"}
	c.JSON(http.StatusOK, resp)
}
