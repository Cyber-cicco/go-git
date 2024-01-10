package cmd

import (
	"fr/eduprolo/src/services"

	"github.com/spf13/cobra"
)

var file = ""
var fileType = "blob"

var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Permet de visualiser le contenu décompressé d'un objet git",
	Long: `
    Permet d'initialiser un nouveau repo git dans le dossier courant.
  `,

	Run: func(cmd *cobra.Command, args []string) {
        services.CatFile(file)
    },
}

func init() {
    catCmd.Flags().StringVarP(&fileType, "type", "t", "blob", "Type du fichier que l'on souhaite désérialiser")
    catCmd.Flags().StringVarP(&file, "file", "f", "", "Nom du fichier que l'on souhaite désérialiser")
    catCmd.MarkFlagRequired("file")
}
