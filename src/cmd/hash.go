package cmd

import (
	"fr/eduprolo/src/services"

	"github.com/spf13/cobra"
)


var objectToHash = ""

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "permet de hasher le contenu d'un fichier",
	Long: `
Permet de visualiser le contenu décompressé d'un objet git
  `,

	Run: func(cmd *cobra.Command, args []string) {
        services.HashFile(objectToHash)
    },
}

func init() {
    hashCmd.Flags().StringVarP(&objectToHash, "file", "f", "", "Nom du fichier que l'on souhaite sérialiser")
    hashCmd.MarkFlagRequired("file")
}
