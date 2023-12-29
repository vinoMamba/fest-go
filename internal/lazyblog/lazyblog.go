package lazyblog

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cfgFile string

func NewLazyBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "lazyblog",
		Short:        "lazyblog App",
		Long:         "lazyblog is a static site generator written in Go",
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数，函数执行失败会返回错误信息
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
		// 这里设置命令运行时，不需要指定命令行参数
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		},
	}
	cobra.OnInitialize(initConfig)
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the lazyblog configuration file. Empty string for no configuration file.")
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	return cmd
}

// run 函数是实际的业务代码入口函数.
func run() error {
	fmt.Println("Hello MiniBlog!")
	return nil
}
