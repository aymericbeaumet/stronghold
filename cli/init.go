package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/aymericbeaumet/stronghold/stronghold"
)

func init() {
	RootCmd.AddCommand(InitMainCmd)
}

// InitMainCmd takes care of initializing the repository.
var InitMainCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize stronghold",
	Long: `Initialize stronghold. The command starts from the current directory
and walks up until it finds either a Stronghold or a VCS marker. If no marker is
found, it falls back to the current directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := stronghold.UpsertStrongholdDirectory()
		if err != nil {
			fmt.Println("Failed", err)
			return
		}
	},
}
