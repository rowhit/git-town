package cmd

import (
	"github.com/Originate/git-town/src/cfmt"
	"github.com/Originate/git-town/src/git"
	"github.com/Originate/git-town/src/stringtools"
	"github.com/spf13/cobra"
)

var offlineCommand = &cobra.Command{
	Use:   "offline [(true | false)]",
	Short: "Displays or sets offline mode",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			printOfflineFlag()
		} else {
			setOfflineFlag(stringtools.StringToBool(args[0]))
		}
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			err := validateBooleanArgument(args[0])
			if err != nil {
				return err
			}
		}
		return validateMaxArgs(args, 1)
	},
}

func printOfflineFlag() {
	cfmt.Println(git.GetPrintableOfflineFlag())
}

func setOfflineFlag(value bool) {
	git.UpdateOffline(value)
}

func init() {
	RootCmd.AddCommand(offlineCommand)
}
