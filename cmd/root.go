package cmd

import (
	"github.com/XuHandsome/stocks/pkgs"
	"github.com/XuHandsome/stocks/pkgs/config"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "stocks",
	Short: "Stocks Command Client",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// 获取配置
		configPath, _ := cmd.Flags().GetString("config")
		mainConfig := config.InitConf(configPath)

		//pkgs.TestView()
		err := pkgs.Run(mainConfig)
		if err != nil {
			return
		}
	},
}

func completionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "completion",
		Short: "Generate the autocompletion script for the specified shell",
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	// 隐藏 completion
	completion := completionCommand()
	completion.Hidden = true
	rootCmd.AddCommand(completion)

	// 参数
	rootCmd.Flags().StringP("config", "", "./config.yaml", "configure file")
}
