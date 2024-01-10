package cmd

import (
	"fr/eduprolo/src/services"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
    Use:   "test",
    Short: "Commande de test",
    Long: `Commande de test `,
    Run: func(cmd *cobra.Command, args []string) {
        services.TestWriteFile("main.go")
    },
}

func init() {
}
