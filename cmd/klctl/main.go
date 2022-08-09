package main

import (
	"context"
	"fmt"
	"os"

	"github.com/kunlun-qilian/klctl/cmd/klctl/gen"
	"github.com/kunlun-qilian/klctl/version"
	"github.com/spf13/cobra"
)

var cmdRoot = &cobra.Command{
	Use:     "klctl",
	Version: version.Version,
}

func init() {
	cmdRoot.AddCommand(gen.CmdGen)
}

func main() {
	if err := cmdRoot.ExecuteContext(context.Background()); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
