package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/aymericbeaumet/stronghold/stronghold"
)

func init() {
	RootCmd.AddCommand(DeinitMainCmd)
}

// DeinitMainCmd takes care of initializing the repository.
var DeinitMainCmd = &cobra.Command{
	Use:   "deinit",
	Short: "Deinitialize stronghold",
	Long:  `Remove any trace of Stronghold.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := stronghold.WipeStrongholdDirectory()
		if err != nil {
			fmt.Println("Failed", err)
			return
		}
	},
}
