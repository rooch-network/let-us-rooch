// Package middlewares 存放系统中间件
package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gohub/pkg/logger"
	"gohub/pkg/response"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

type LogFields struct {
	Code         int `json:"code,omitempty"`
	Request      string
	RequestBody  string
	ResponseBody string
	IP           string
	Status       int
	UserAgent    string
	Errors       string
	Time         string
}

// Logger 记录请求日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取 response 内容
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		// 获取请求数据
		var requestBody []byte
		if c.Request.Body != nil {
			// c.Request.Body 是一个 buffer 对象，只能读取一次
			requestBody, _ = io.ReadAll(c.Request.Body)
			// 读取后，重新赋值 c.Request.Body ，以供后续的其他操作
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 设置开始时间
		start := time.Now()
		c.Next()

		// 开始记录日志的逻辑
		cost := time.Since(start)
		responseStatus := c.Writer.Status()

		// 响应的内容
		responseBody := w.body.Bytes()
		logFields := LogFields{
			Status:       responseStatus,
			Request:      c.Request.Method + " " + c.Request.URL.String(),
			RequestBody:  string(requestBody),
			ResponseBody: string(responseBody),
			IP:           c.ClientIP(),
			UserAgent:    c.Request.UserAgent(),
			Errors:       c.Errors.ByType(gin.ErrorTypePrivate).String(),
			Time:         fmt.Sprintf("%.3fms", float64(cost.Nanoseconds())/1e6),
		}
		if responseStatus == 200 {
			contentType := c.Writer.Header().Get("Content-Type")
			if contentType == "application/json" {
				// 尝试将响应体转换为 CommonResult
				var commonResult response.CommonResult
				err := json.Unmarshal(w.body.Bytes(), &commonResult)
				if err == nil {
					logFields.Code = commonResult.Code
					if commonResult.Code == 200 {
						logger.Debug("HTTP Request:", logFields)
					} else {
						logger.Warn("HTTP Request:", logFields)
					}
				} else {
					logger.Errorv(err)
				}
			}
		} else {
			logger.Warn("HTTP Request:", logFields)
		}
	}
}
