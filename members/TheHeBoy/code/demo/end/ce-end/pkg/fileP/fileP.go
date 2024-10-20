package fileP

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gohub/pkg/app"
	"gohub/pkg/config"
	"os"
	"strings"
)

func GetRootPath() string {
	basePath := config.Get("file_base_path")
	return fmt.Sprintf("%s/public/", strings.TrimSuffix(basePath, "/"))
}

func CombinePath(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return GetRootPath() + path
}

func Sava(business string, data []byte, filename string) (string, error) {
	// 确保目录存在，不存在创建
	filePath := fmt.Sprintf("/%s/%s/", business, app.TimenowInTimezone().Format("2006/01/02"))
	savePath := CombinePath(filePath)
	err := os.MkdirAll(savePath, 0655)
	if err != nil {
		return "", errors.WithStack(err)
	}

	filename = uuid.New().String() + "-" + filename
	// 保存文件
	err = os.WriteFile(savePath+filename, data, 0644)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return filePath + filename, nil
}
