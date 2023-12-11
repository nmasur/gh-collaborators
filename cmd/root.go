package cmd

import (
	"github.com/spf13/cobra"

	userCmd "github.com/katiem0/gh-collaborators/cmd/user"
)

func NewCmdRoot() *cobra.Command {

	cmdRoot := &cobra.Command{
		Use:   "collaborators <command> [flags]",
		Short: "List guest collaborators and their repos.",
		Long:  "List guest collaborators and their repos.",
	}

	//cmdRoot.AddCommand(listCmd.NewCmdList())
	cmdRoot.AddCommand(userCmd.NewCmdList())
	cmdRoot.CompletionOptions.DisableDefaultCmd = true
	cmdRoot.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
	return cmdRoot
}
