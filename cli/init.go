package cli

import (
	"errors"
	"os"

	"github.com/spf13/cobra"

	"github.com/aymericbeaumet/stronghold/lib/pgp"
	"github.com/aymericbeaumet/stronghold/lib/stronghold"
	strongholdGroup "github.com/aymericbeaumet/stronghold/lib/stronghold/group"
	strongholdUser "github.com/aymericbeaumet/stronghold/lib/stronghold/user"
)

// InitMainCmd takes care of initializing the repository.
var InitMainCmd = &cobra.Command{
	Use:   "init identifier",
	Short: "Initialize stronghold",
	Long: `Initialize stronghold. The command starts from the current directory
and walks up until it finds either a Stronghold or a VCS marker. If no marker is
found, it falls back to the current directory.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("The private key identifier is mandatory")
		}
		_, err := pgp.New(pgp.NewOptions{
			Identifier: args[0],
			Gpg:        Gpg,
			Homedir:    Homedir,
		})
		if err != nil {
			return err
		}
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		s, err := stronghold.New(stronghold.NewOptions{
			WorkingDirectory: wd,
		})
		if err != nil {
			return err
		}
		g, err := s.CreateGroup(strongholdGroup.NewOptions{
			Name: "root",
		})
		if err != nil {
			return err
		}
		if _, err := s.AddUser(strongholdUser.NewOptions{
			Identifier: args[0],
			Group:      g,
		}); err != nil {
			return err
		}
		if err = s.Dump(false); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(InitMainCmd)
}
