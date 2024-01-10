package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "git",
  Short: "Tentative de recréation des fonctionalités de base de git",
  Long: 
`Tentative de recréation des fonctionalités de base de git
`,
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}

func init() {
    rootCmd.AddCommand(addCmd)
    rootCmd.AddCommand(initCmd)
    rootCmd.AddCommand(testCmd)
    rootCmd.AddCommand(catCmd)
}
