package cmd

import "github.com/spf13/cobra"

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Init collection (create)",
	Run: func(cmd *cobra.Command, args []string) {
		/*entityDir, err := storage.Path()
		if err != nil {
			fmt.Println("could not get the root storage directory:", err)
			return
		}

		// TODO: Maybe display that it is downloading
		err = entity.Get(entityDir)
		if err != nil {
			fmt.Println("Error crawling directories:", err)
			return
		}*/
	},
}
