package services

import (
	"fr/eduprolo/src/consts"
	"fr/eduprolo/src/entites"
	"fr/eduprolo/src/ioutils"
	"fr/eduprolo/src/utils"
	"os"
)

func AddFile(file string) {
    wd, err := os.Getwd()
    utils.HandleBasicError(err, consts.ERR_CANT_READ_DIR)
    rootRepositoryPath := ioutils.FindRootRepository(wd)
    repo := entites.GetGitRepoFromPath(rootRepositoryPath)
    ioutils.ReadObject(repo, "055c2eb5bc0ec5bc3340dcf6044cf8476dc8e154")
}

