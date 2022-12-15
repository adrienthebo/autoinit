package autoinit

import (
	"fmt"

	_ "embed"

	"github.com/spf13/cobra"
)

//go:embed init.bash
var initBash string

func InitCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize autoinit in your shell",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(initBash)
		},
	}
}
