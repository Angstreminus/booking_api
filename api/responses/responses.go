package response

import (
	"github.com/gin-gonic/gin"
)

type ErrorData struct {
	Error  string
	Params interface{}
	Method string
	Path   string
	Host   string
	Header interface{}
}

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(code, Response{Data: nil, Message: message})
}

func Json(c *gin.Context, code int, data interface{}) {
	c.JSON(code, Response{Data: data, Message: "OK"})
}
