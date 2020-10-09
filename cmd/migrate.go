package cmd

import (
	"github.com/fahmisyaifudin/echo-boilerplate/database"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	Long:  `Migrate the database`,
	Run: func(cmd *cobra.Command, args []string) {
		database.Migrate()
	},
}
