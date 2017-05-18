package cli

import "github.com/spf13/cobra"

// RootCmd is the first command of this program, from which all the other
// commands will inherit.
var RootCmd = &cobra.Command{
	Use:   "stronghold",
	Short: "Stronghold allows you to safely store encrypted secrets",
	Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// Gpg contains the GPG binary to be used (e.g.: gpg, gpg1 or gpg2)
var Gpg = ""

// Homedir contains the GPG home directory (e.g.: ~/.gnupg)
var Homedir = ""

func init() {
	RootCmd.PersistentFlags().StringVar(&Gpg, "gpg", "gpg", "gpg binary")
	RootCmd.PersistentFlags().StringVar(&Homedir, "homedir", "", "gpg home directory")
}
