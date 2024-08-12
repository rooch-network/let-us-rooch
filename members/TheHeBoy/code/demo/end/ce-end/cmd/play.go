package cmd

import (
	"github.com/spf13/cobra"
	"gohub/pkg/logger"
	"gohub/pkg/rooch"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

// 调试完成后请记得清除测试代码
func runPlay(cmd *cobra.Command, args []string) {
	//data, err := rooch.BtcQueryInscriptions("tb1p588cfjvg4d6hchnuslzx6zhchzq3ht4csqmfeafhvqn5g680c74s89day6")
	data, err := rooch.IsOpens()
	if err != nil {
		logger.Error(err)
	}

	logger.Info(data)
}
