// Package response 响应处理工具
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommonResult struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, CommonResult{
		Code: 200,
		Data: nil,
		Msg:  "success",
	})
}

func SuccessData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, CommonResult{
		Code: 200,
		Data: data,
		Msg:  "success",
	})
}

func Error(c *gin.Context, err error) {
	c.JSON(http.StatusOK, CommonResult{
		Code: 500,
		Data: nil,
		Msg:  err.Error(),
	})
}

func ErrorStr(c *gin.Context, err string) {
	c.JSON(http.StatusOK, CommonResult{
		Code: 500,
		Data: nil,
		Msg:  err,
	})
}

func errorCode(c *gin.Context, code int, err error) {
	c.JSON(http.StatusOK, CommonResult{
		Code: code,
		Data: nil,
		Msg:  err.Error(),
	})
}

// Error405 参数校验错误
func Error405(c *gin.Context, err error) {
	errorCode(c, 405, err)
}
