package cli

import "github.com/spf13/cobra"

func init() {
}

// RootCmd is the first command of this program, from which all the other
// commands will inherit.
var RootCmd = &cobra.Command{
	Use:   "stronghold",
	Short: "Stronghold allows you to safely store encrypted secrets",
	Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}
