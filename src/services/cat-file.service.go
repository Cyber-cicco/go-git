package services

import (
	"fmt"
	"fr/eduprolo/src/ioutils"
)

func CatFile(file string) {
    if(!ioutils.FileExists(file)) {
        fmt.Println("Le fichier n'existe pas")
        panic("Erreur : le fichier n'existe pas")
    }
    gitObject := ReadObject(file)
    content := gitObject.Deserialize()
    fmt.Println(content)
}
