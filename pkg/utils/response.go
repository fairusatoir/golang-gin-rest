package utils

import "github.com/gin-gonic/gin"

type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, BaseResponse{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func NewErrorResponse(c *gin.Context, statusCode int, err string) {
	c.JSON(statusCode, BaseResponse{
		Status:  false,
		Message: err,
	})

}