package cmd

import (
	"github.com/akerl/syncgithub/repos"
	"github.com/spf13/cobra"
)

func updateRunner(cmd *cobra.Command, _ []string) error {
	c, err := cmd.Flags().GetString("config")
	if err != nil {
		return err
	}

	config, err := repos.LoadConfig(c)
	if err != nil {
		return err
	}

	r, err := repos.FilteredRepos(config)
	if err != nil {
		return err
	}

	return repos.HandleRepos(config, r)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update local copy of repos",
	RunE:  updateRunner,
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("config", "c", "", "Config file")
}
