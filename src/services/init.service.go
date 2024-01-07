package services

import (
	"encoding/json"
	"fr/eduprolo/src/consts"
	"fr/eduprolo/src/entites"
	"fr/eduprolo/src/ioutils"
	"fr/eduprolo/src/utils"
	"os"
)

func ExecuteInit(path string) {

    currDir, err := os.Getwd();
    utils.HandleBasicError(err, consts.ERR_CANT_READ_DIR)
    gitDir := currDir + "/.git/"

    if ioutils.FileExists(gitDir) {
        println("Il existe déjà un dossier .git dans le dossier courant.")
        os.Exit(1)
    }

    ioutils.CreateDirectoryAtPath(gitDir)
    ioutils.CreateDirectoryAtPath(gitDir + consts.DIR_OBJECTS)
    ioutils.CreateDirectoryAtPath(gitDir + consts.DIR_REFS)
    ioutils.CreateDirectoryAtPath(gitDir + consts.DIR_INFOS)
    ioutils.CreateDirectoryAtPath(gitDir + consts.DIR_BRANCHES)
    ioutils.CreateDirectoryAtPath(gitDir + consts.DIRE_REFS_HEADS)
    ioutils.CreateDirectoryAtPath(gitDir + consts.DIRE_REFS_TAGS)

    head := entites.GetBaseHead()
    config := entites.GetBaseConfigFile()

    headString, err := json.MarshalIndent(head, "", "  ")
    utils.HandleBasicError(err, consts.ERR_CANT_CREATE_DIR)

    configString, err := json.MarshalIndent(config, "", "  ")
    utils.HandleBasicError(err, consts.ERR_CANT_CREATE_DIR)

    ioutils.WriteFileAtPath([]byte(headString), gitDir + consts.FILE_HEAD)
    ioutils.WriteFileAtPath([]byte(configString), gitDir + consts.FILE_CONFIG)
    ioutils.WriteFileAtPath([]byte("Unnamed repository; edit this file 'description' to name the repository."), gitDir + consts.FILE_DESCRIPTION)

}

