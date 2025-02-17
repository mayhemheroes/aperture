package blueprints

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	apertureBlueprintsRepoDefault    = "github.com/fluxninja/aperture/blueprints"
	apertureBlueprintsVersionDefault = "main"
)

var (
	blueprintsDir string

	// Args for `blueprints`.
	blueprintsVersion     string
	apertureBlueprintsURI string
)

func init() {
	BlueprintsCmd.PersistentFlags().StringVar(&blueprintsVersion, "version", apertureBlueprintsVersionDefault, "version of aperture blueprint")
	BlueprintsCmd.PersistentFlags().StringVar(&apertureBlueprintsURI, "uri", apertureBlueprintsRepoDefault, "URI of aperture blueprints, could be a local path or a remote git repository")

	BlueprintsCmd.AddCommand(pullCmd)
	BlueprintsCmd.AddCommand(listCmd)
	BlueprintsCmd.AddCommand(removeCmd)
	BlueprintsCmd.AddCommand(generateCmd)
}

// BlueprintsCmd is the root command for blueprints.
var BlueprintsCmd = &cobra.Command{
	Use:   "blueprints",
	Short: "Manage blueprints",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		blueprintsDir = filepath.Join(userHomeDir, ".aperturectl", "blueprints")
		err = os.MkdirAll(blueprintsDir, os.ModePerm)
		if err != nil {
			return err
		}
		if apertureBlueprintsURI == "" {
			apertureBlueprintsURI = apertureBlueprintsRepoDefault
		}
		return nil
	},
}
