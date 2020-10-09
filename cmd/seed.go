package cmd

import (
	"github.com/fahmisyaifudin/echo-boilerplate/database"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(seedCmd)
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Run Seeder the database",
	Long:  `Run Seeder the database`,
	Run: func(cmd *cobra.Command, args []string) {
		database.RunSeeder()
	},
}
