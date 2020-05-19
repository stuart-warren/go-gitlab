package create

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/makkes/gitlab-cli/api"
	"github.com/spf13/cobra"
)

func createCommand(client api.Client, args []string, issueDetails []string, out io.Writer) error {
	project, err := client.FindProject(args[0])
	if err != nil {
		return err
	}
	var issueTitle, issueDescription string
	if len(issueDetails) == 0 {
		// FIXME: add user input (git commit like) including template
		return fmt.Errorf("Currently you must provide a message flag")
	}
	issueTitle = issueDetails[0]
	if len(issueDetails) == 1 {
		issueDescription = ""
	} else {
		issueDescription = strings.Join(issueDetails[1:], "\n\n")
	}

	resp, _, err := client.Post("/projects/"+strconv.Itoa(project.ID)+"/issues", strings.NewReader(fmt.Sprintf("title=%s&description=%s", url.QueryEscape(issueTitle), url.QueryEscape(issueDescription))))
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	json.Indent(&buf, resp, "", "    ")
	buf.WriteTo(out)
	out.Write([]byte("\n"))
	return nil

}

func NewCommand(client api.Client) *cobra.Command {
	var issueDetails []string
	cmd := &cobra.Command{
		Use:   "create new issue",
		Short: "Provide detailed information for a new issue",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return createCommand(client, args, issueDetails, os.Stdout)
		},
	}
	cmd.Flags().StringSliceVarP(&issueDetails, "message", "m", []string{}, "Use the given <msg> as the issue title. If multiple -m options are given, their values are concatenated as separate paragraphs in the description.")
	cmd.MarkFlagRequired("message") // FIXME: related to above FIXME
	return cmd
}
