package cli

import (
	"errors"
	"os"

	"github.com/spf13/cobra"

	"github.com/aymericbeaumet/stronghold/lib/crypto"
	"github.com/aymericbeaumet/stronghold/lib/stronghold"
)

// InitCommand takes care of initializing the repository.
var InitCommand = &cobra.Command{
	Use:   "init identifier",
	Short: "Initialize stronghold",
	Long: `Initialize stronghold. The command starts from the current directory
and walks up until it finds either a Stronghold or a VCS marker. If no marker is
found, it falls back to the current directory.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Initialize the key box with the private key
		if len(args) < 1 {
			return errors.New("The private key identifier is mandatory")
		}
		keyBox := crypto.NewKeyBox().SetPrivateKey(crypto.NewPrivateKey(args[0], Gpg, Homedir))
		// Initialize the safe (walking up from the working directory), then commit
		// to the disk
		workingDirectory, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		safe := stronghold.NewVault(workingDirectory, keyBox)
		user := stronghold.NewUser(args[0])
		group := stronghold.NewGroup("root").AddUser(user)
		safe.AddGroup(group)
		safe.AddUser(user)
		return safe.Commit()
	},
}

func init() {
	RootCommand.AddCommand(InitCommand)
}
