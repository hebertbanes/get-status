package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIBadRequestMsg : helper for c.JSON response
func APIBadRequestMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": true, "msg": msg})
}

// APIBadRequestData : helper for c.JSON response
func APIBadRequestData(c *gin.Context, data map[string]interface{}) {
	result := map[string]interface{}{
		"error": true,
		"msg":   "error",
	}
	for k, v := range data {
		result[k] = v
	}
	c.JSON(http.StatusBadRequest, result)
}

// APIOkMsg : helper for c.JSON response
func APIOkMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{"error": true, "msg": msg})
}

// APIOkData : helper for c.JSON response
func APIOkData(c *gin.Context, data interface{}) {
	result := map[string]interface{}{
		"error": false,
		"msg":   "ok",
		"data":  data,
	}
	c.JSON(http.StatusOK, result)
}

// APIOkDataMsg : helper for c.JSON response
func APIOkDataMsg(c *gin.Context, data interface{}, msg string) {
	result := map[string]interface{}{
		"error": false,
		"msg":   msg,
		"data":  data,
	}
	c.JSON(http.StatusOK, result)
}
