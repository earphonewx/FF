package cmd

import (
	"ff/initialize"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ff",
	Short: "Hello, ff-cli will serve for you!",
	Long: "A simple system which contains some functions for daily use. \n\t Just for fun!",
	// 预加载配置、初始化logger, 持续预运行
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 初始化配置文件
		initialize.InitViper()
		// 初始化日志记录器
		initialize.InitLogger()
	},
	Run: func(cmd *cobra.Command, args []string) {
		// do something here.
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	//rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")

	//rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	//rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	//viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	//viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	//viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	//viper.SetDefault("license", "apache")

	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(webServerCmd)
}