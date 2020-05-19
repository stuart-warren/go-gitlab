package issue

import (
	"github.com/makkes/gitlab-cli/api"
	"github.com/makkes/gitlab-cli/cmd/issue/create"
	"github.com/makkes/gitlab-cli/cmd/issue/inspect"
	"github.com/spf13/cobra"
)

func NewCommand(client api.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issue COMMAND",
		Short: "Manage issues",
	}

	cmd.AddCommand(create.NewCommand(client))
	cmd.AddCommand(inspect.NewCommand(client))

	return cmd
}
