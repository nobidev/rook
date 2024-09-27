package main

import (
	"time"

	rook "github.com/rook/rook/cmd/rook/rook"
	"github.com/rook/rook/pkg/daemon/discover"
	"github.com/rook/rook/pkg/util/flags"
	"github.com/spf13/cobra"
)

var (
	discoverTestCmd = &cobra.Command{
		Use: "discover-test",
	}
)

func init() {
	flags.SetFlagsFromEnv(discoverTestCmd.Flags(), rook.RookEnvVarPrefix)
	discoverTestCmd.RunE = startDiscoverTest
}

func startDiscoverTest(cmd *cobra.Command, args []string) error {
	rook.SetLogLevel()

	context := rook.NewContext()
	ctx := cmd.Context()

	err := discover.Run(ctx, context, time.Minute, true)
	if err != nil {
		rook.TerminateFatal(err)
	}

	return nil
}
