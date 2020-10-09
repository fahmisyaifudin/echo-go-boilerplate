package cmd

import (
	"github.com/fahmisyaifudin/echo-boilerplate/database"
	"github.com/fahmisyaifudin/echo-boilerplate/route"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start Server",
	Long:  `Start Server`,
	Run: func(cmd *cobra.Command, args []string) {
		db := database.Connect()
		route.HandleRequest(db)
	},
}
