package main

import (
	"building_management/cmd"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "Building Management System",
		Short: "Building Management System api",
	}
	rootCmd.AddCommand(
		cmd.API(),
		cmd.Migrations(),
	)
	cobra.CheckErr(rootCmd.Execute())
}
