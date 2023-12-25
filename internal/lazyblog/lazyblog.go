package lazyblog

import "github.com/spf13/cobra"

func NewLazyBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lazyblog",
		Short: "lazyblog is a static site generator written in Go",
		Long:  "lazyblog is a static site generator written in Go",
	}
	return cmd
}
