package gen

import (
	"fmt"

	"github.com/kunlun-qilian/klctl/internal/generate"
	"github.com/spf13/cobra"
)

var (
	cmdGenClientFlagSpecURL              string
	cmdGenClientFlagOpenClientOutputPath string
)

func init() {
	CmdGen.AddCommand(cmdGenClient)

	cmdGenClient.Flags().
		StringVarP(&cmdGenClientFlagSpecURL, "spec-url", "", "", "client spec url")
	cmdGenClient.Flags().
		StringVarP(&cmdGenClientFlagOpenClientOutputPath, "output", "", ".", "openapi client spec file path")
}

var cmdGenClient = &cobra.Command{
	Use:     "client",
	Example: "client demo",
	Short:   "generate client by open api",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			panic(fmt.Errorf("need service name"))
		}

		if cmdGenClientFlagSpecURL != "" {
			cmdGenClientFlagSpecURL += "?format=yaml"
		}

		generate.GenerateOpenapiClient(args[0], cmdGenClientFlagOpenClientOutputPath, cmdGenClientFlagSpecURL)

	},
}
