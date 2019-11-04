package httpHandlers

import (
	"github.com/gin-gonic/gin"
)

// health-check
// /health
func Health(c *gin.Context) {
	resp := map[string]string{"status": "ok"}
	c.JSON(200, resp)
	//c.Set(config.KeyResponse, resp)
	//c.JSON(http.StatusOK, map[string]string{"error": err.Error()})

	c.Abort()
}
