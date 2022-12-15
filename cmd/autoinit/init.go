package autoinit

import (
	"fmt"

	"github.com/spf13/cobra"
)

func InitCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize autoinit in your shell",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello!")
		},
	}
}
