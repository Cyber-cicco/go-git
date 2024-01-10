package services

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"fr/eduprolo/src/consts"
	"fr/eduprolo/src/entites"
	"fr/eduprolo/src/ioutils"
	"fr/eduprolo/src/utils"
	"io"
	"os"
	"strconv"
)

func getRootRepo() entites.GitRepository {
    wd, err := os.Getwd()
    utils.HandleBasicError(err, consts.ERR_CANT_READ_DIR)
    rootRepositoryPath := ioutils.FindRootRepository(wd)
    return entites.GetGitRepoFromPath(rootRepositoryPath)
}

func ReadObject(path string) entites.GitObject {
    f := *bytes.NewBuffer(ioutils.ReadFile(path))

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
    switch format {
    case "blob": {
        return &entites.GitBlob {
            Content: content,
            Length: int(length),
        }
    }
    case "flag" :  {
        return &entites.GitFlag{}
    }
    case "commit" : {
        return &entites.GitCommit{}
    }
    case "tree" : {
        return &entites.GitTree{}
    }
    }
    panic("Format invalide")
}

func TestWriteFile(file string) {
    repo := getRootRepo()
    writeObject(repo, file)
}

func AddFile(file string) {

}

func writeObject(repo entites.GitRepository, file string) {
    fileContent := ioutils.ReadFile(file)
    fileLen := len(fileContent)
    fmt.Println(fileLen)
    gitContent := fmt.Sprintf("blob %d%c%s", fileLen, '\x00', fileContent)
    sha :=  sha1.Sum([]byte(gitContent))
    name := base64.RawURLEncoding.EncodeToString(sha[:])
    dirName := string(name[:2]) + "/"
    fileName := string(name[2:])
    var b bytes.Buffer
    w := zlib.NewWriter(&b)
    w.Write([]byte(gitContent))
    w.Close()
    compressedContent := b.Bytes()
    fmt.Println(compressedContent)
    ioutils.WriteSimpleFile(dirName, fileName, compressedContent)
}

