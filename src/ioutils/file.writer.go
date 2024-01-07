package ioutils

import (
	"fr/eduprolo/src/consts"
	"fr/eduprolo/src/utils"
	"os"
)

func CreateDirectoryAtPath(path string) {
    err := os.MkdirAll(path, 0777)
    utils.HandleBasicError(err, consts.ERR_CANT_CREATE_DIR)
}
    
func WriteFileAtPath(bytes []byte, filename string) {
    if FileExists(filename) {
        overrideFile(bytes, filename)
    } else {
        createNewFile(bytes, filename)
    }
}

func createNewFile(bytes []byte, filename string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0660)
    defer f.Close()
	utils.HandleBasicError(err, consts.ERR_CANT_CREATE_FILE)
	f.Write(bytes)
	f.Sync()
}

func overrideFile(bytes []byte, filename string) {
    f, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
    defer f.Close()
	utils.HandleBasicError(err, consts.ERR_CANT_CREATE_FILE)
    f.Write(bytes)
    f.Sync()
}

func FileExists(dirPath string) bool {
	_, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
        panic(err)
	}
	return true
}
