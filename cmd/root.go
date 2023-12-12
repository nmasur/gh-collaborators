package cmd

import (
	"github.com/spf13/cobra"

	addCmd "github.com/katiem0/gh-collaborators/cmd/add"
	listCmd "github.com/katiem0/gh-collaborators/cmd/list"
	removeCmd "github.com/katiem0/gh-collaborators/cmd/remove"
)

func NewCmdRoot() *cobra.Command {

	cmdRoot := &cobra.Command{
		Use:   "collaborators <command> [flags]",
		Short: "List and maintain repository collaborators and their repos.",
		Long:  "List and maintain repository collaborators and their assigned repositories.",
	}

	cmdRoot.AddCommand(addCmd.NewCmdAdd())
	cmdRoot.AddCommand(listCmd.NewCmdList())
	cmdRoot.AddCommand(removeCmd.NewCmdRemove())
	cmdRoot.CompletionOptions.DisableDefaultCmd = true
	cmdRoot.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
	return cmdRoot
}
