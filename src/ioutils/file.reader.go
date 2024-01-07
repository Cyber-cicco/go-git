package ioutils

import (
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"fr/eduprolo/src/consts"
	"fr/eduprolo/src/entites"
	"fr/eduprolo/src/utils"
	"io"
	"os"
	"strconv"
)

func FindRootRepository(path string) string {

    if path == "" {
        println("Aucun dépôt git n'a été initialisé dans le répertoire courant. Utiliser git init pour remédier à cela")
        os.Exit(1)
    }

    if !FileExists(path) {
        println("Le répertoire est censé exister")
        os.Exit(1)
    }

    dir, err := os.ReadDir(path)

    utils.HandleBasicError(err, consts.ERR_CANT_READ_DIR)
    
    for _, file := range dir {
        if file.IsDir() && file.Name() == ".git" {
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

func ReadObject(repo entites.GitRepository, sha string) {
    f := repo.GetObjectFileFromHash(sha);

    reader, err := zlib.NewReader(&f)
    utils.HandleBasicError(err, consts.ERR_CANT_DECOMPRESS)

    var raw bytes.Buffer
    io.Copy(&raw, reader)

    contentAsBytes := raw.Bytes()
    iFormatSeparator := utils.IndexOf(raw.Bytes(), ' ')
    iLenSeparator := utils.IndexOf(raw.Bytes(), '\x00')
    format := string(contentAsBytes[:iFormatSeparator])
    length, err := strconv.ParseInt(string(contentAsBytes[iFormatSeparator+1:iLenSeparator]), 10, 0)
    utils.HandleBasicError(err, consts.ERR_GIT_FILE_MALFORMED)
    content := string(contentAsBytes[iLenSeparator+1:])

    if len(content) != int(length) {
        utils.HandleBasicError(errors.New("Erreur dans la longueur du fichier"), consts.ERR_GIT_FILE_MALFORMED)
    }

    fmt.Println(format)
    fmt.Println(length)
    fmt.Println(content)
}
