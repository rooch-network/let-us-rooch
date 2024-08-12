package validators

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/thedevsaddam/govalidator"
	"gohub/pkg/logger"
	"gohub/pkg/response"
)

type ValidatorI interface {
	Validator() map[string][]string
}

func Validate(c *gin.Context, valObj ValidatorI) bool {

	// 1. 解析请求，支持 JSON 数据、表单请求和 URL Query
	if err := c.ShouldBind(valObj); err != nil {
		logger.Errorv(err)
		response.Error405(c, errors.New("请求解析错误，请确认请求格式是否正确。参数请使用 JSON 格式。"))
		return false
	}

	// 2. 表单验证
	errs := valObj.Validator()

	jsonData, err := json.Marshal(errs)
	if err != nil {
		response.Error(c, err)
		return false
	}

	// 3. 判断验证是否通过
	if len(errs) > 0 {
		response.Error405(c, errors.New(fmt.Sprintf("参数校验错误: %s", string(jsonData))))
		return false
	}

	return true
}

func ValidateData(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	// 配置选项
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}
	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}
