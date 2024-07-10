package cmd

import (
	"fmt"
	"github.com/AmadlaOrg/hery/env"
	"github.com/AmadlaOrg/hery/storage"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
)

var SettingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "List the paths and other environment variables for HERY",
	Run: func(cmd *cobra.Command, args []string) {
		heryPath, err := storage.Path()
		if err != nil {
			fmt.Println(err)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Setting", "Value"})
		table.Append([]string{"Collections path", heryPath})

		for _, varName := range env.List() {
			val := os.Getenv(varName)
			table.Append([]string{varName, val})
		}

		table.Render()
	},
}
