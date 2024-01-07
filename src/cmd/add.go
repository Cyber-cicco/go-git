package cmd

import (
	"fr/eduprolo/src/services"

	"github.com/spf13/cobra"
)

var fileName string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Ajouter un fichier",
	Long: `
    Ajouter un fichier dans la partie staged de l'application
  `,

	Run: func(cmd *cobra.Command, args []string) {
        services.AddFile(fileName)
    },
}

func init() {
    addCmd.Flags().StringVarP(&fileName, "file", "f", "", "Fichier que vous souhaitez ajouter")
    addCmd.MarkFlagRequired("file")
}
