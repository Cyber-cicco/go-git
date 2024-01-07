package ioutils

import (
	"fr/eduprolo/src/consts"
	"fr/eduprolo/src/utils"
	"os"
)

func FindRootRepository(path string) string {

    if !FileExists(path) {
        println("Le répertoire est censé exister")
        os.Exit(1)
    }

    if path == "" {
        println("Aucun dépôt git n'a été initialisé dans le répertoire courant. Utiliser git init pour remédier à cela")
        os.Exit(1)
    }

    dir, err := os.ReadDir(path)

    utils.HandleBasicError(err, consts.ERR_CANT_READ_DIR)
    
    for _, file := range dir {
        if file.IsDir() && file.Name() == ".git/" {
            return path;
        }
    }

    return FindRootRepository(walkToParentDirectory(path))
}

func walkToParentDirectory(path string) string {
    indexOfLastSlash := 0;
    i := 0
    for i < len(path) {
        if path[i] == '\\' {
            i += 2
            continue
        }
        if path[i] == '/' {
            indexOfLastSlash = i
        }
        i += 1
    }
    return path[:indexOfLastSlash]
}
