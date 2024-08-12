// Package config 负责配置信息
package config

import (
	"fmt"
	"os"

	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper" // 自定义包名，避免与内置 viper 实例冲突
)

// viper 库实例
var viper *viperlib.Viper

// InitConfig 完成环境变量
func InitConfig(env string) {
	// 1. 初始化 Viper 库
	viper = viperlib.New()
	// 2. 配置类型，支持 "json", "toml", "yaml", "yml", "properties",
	//             "props", "prop", "env", "dotenv"
	viper.SetConfigType("yaml")
	// 3. 环境变量配置文件查找的路径，相对于 main.go
	viper.AddConfigPath(".")
	// 4. 设置环境变量前缀，用以区分 Go 的系统环境变量
	viper.SetEnvPrefix("appenv")
	// 5. 设置默认值
	viper.SetConfigName("application.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	configs := viper.AllSettings()
	for k, v := range configs {
		viper.SetDefault(k, v)
	}

	loadEnv(env)
}

func loadEnv(envSuffix string) {
	var envPath string
	// 优先级：参数 > 配置文件 > 默认值
	if len(envSuffix) > 0 {
		filepath := formatEnvStr(envSuffix)
		if _, err := os.Stat(filepath); err == nil {
			envPath = filepath
		}
	} else if len(Get("app.env")) > 0 {
		envSuffix = Get("app.env")
		envPath = formatEnvStr(envSuffix)
	}

	if len(envPath) > 0 {
		viper.SetConfigName(envPath)
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.Set("app.env", envSuffix)
	// 监控配置文件变更时重新加载
	viper.WatchConfig()
}

func formatEnvStr(env string) string {
	return fmt.Sprintf("application-%s.yaml", env)
}

// Get 获取配置项
// 第一个参数 path 允许使用点式获取，如：app.name
// 第二个参数允许传参默认值
func Get(path string, defaultValue ...interface{}) string {
	return GetString(path, defaultValue...)
}

func internalGet(path string, defaultValue ...interface{}) interface{} {
	// config 或者环境变量不存在的情况
	if !viper.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(path)
}

// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

// GetFloat64 获取 float64 类型的配置信息
func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, defaultValue...))
}

// GetUint64 获取 Int64 类型的配置信息
func GetUint64(path string, defaultValue ...interface{}) uint64 {
	return cast.ToUint64(internalGet(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(internalGet(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

// GetStringMapString 获取结构数据
func GetStringMapString(path string) map[string]string {
	return viper.GetStringMapString(path)
}
