// cmd/root.go
package cmd

import (
    "github.com/spf13/cobra"
    "fmt"
)

var rootCmd = &cobra.Command{
    Use:   "gopassgen",
    Short: "Generate, Store and Manage passwords",
    Long:  "A command-line tool to generate random passwords and manage them.",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Use 'gopassgen --help' to see available commands.")
    },
}

func Execute() error {
    return rootCmd.Execute()
}
