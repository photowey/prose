package fasthttp

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Fasthttp = &cobra.Command{
	Use:     "fasthttp",
	Aliases: []string{"fh", "fast"},
	Run: func(c *cobra.Command, rawOpts []string) {
		fmt.Println("hello prose fasthttp")
	},
}
