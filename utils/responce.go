package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response structure
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// JSONResponse sends a JSON response
func JSONResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse sends an error response
func ErrorResponse(c *gin.Context, status int, message string) {
	JSONResponse(c, status, message, nil)
}

// SuccessResponse sends a success response
func SuccessResponse(c *gin.Context, data interface{}) {
	JSONResponse(c, http.StatusOK, "success", data)
}
