package cmd

import (
	"fr/eduprolo/src/services"

	"github.com/spf13/cobra"
)

var initDirectory = "."

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialiser un repo git",
	Long: `
    Permet d'initialiser un nouveau repo git dans le dossier courant.
  `,

	Run: func(cmd *cobra.Command, args []string) {
        services.ExecuteInit(initDirectory)
    },
}

func init() {
    initCmd.Flags().StringVarP(&initDirectory, "directory", "d", ".", "Répertoire où initialiser le repo")
}
