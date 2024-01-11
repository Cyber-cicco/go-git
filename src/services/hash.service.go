package services

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"fr/eduprolo/src/ioutils"
)

func HashFile(file string) {
    print(file)
    if(!ioutils.FileExists(file)) {
        fmt.Println("Le fichier n'existe pas")
        panic("Erreur : le fichier n'existe pas")
    }
    fileContent := ioutils.ReadFile(file)
    fileLen := len(fileContent)
    gitContent := fmt.Sprintf("blob %d%c%s", fileLen, '\x00', fileContent)
    var b bytes.Buffer
    w := zlib.NewWriter(&b)
    w.Write([]byte(gitContent))
    w.Close()
    fmt.Println(string(b.Bytes()))
}
