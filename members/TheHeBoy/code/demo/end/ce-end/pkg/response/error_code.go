package response

import (
	"github.com/gin-gonic/gin"
)

func Error10001(c *gin.Context, err error) {
	errorCode(c, 10001, err)
}

func Error10002(c *gin.Context, err error) {
	errorCode(c, 10002, err)
}

func Error10003(c *gin.Context, err error) {
	errorCode(c, 10003, err)
}
