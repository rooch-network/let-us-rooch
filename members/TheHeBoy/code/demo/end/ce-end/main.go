package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"gohub/cmd"
	"gohub/pkg/btcapi"
	"gohub/pkg/config"
	"gohub/pkg/console"
	"gohub/pkg/database"
	"gohub/pkg/hashidsP"
	"gohub/pkg/logger"
	"gohub/pkg/snowflakeP"
	"os"
)

func main() {
	// 应用的主入口，默认调用 cmd.CmdServe 命令
	var rootCmd = &cobra.Command{
		Use: "Gohub",

		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {

			// 初始化
			config.InitConfig(cmd.Env)
			logger.InitLogger()
			database.InitDB()
			snowflakeP.InitSnowflake()
			btcapi.InitBtc()
			hashidsP.InitHashIds()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdPlay,
	)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
