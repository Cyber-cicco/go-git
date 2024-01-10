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

/* WriteSimpleFile 
*  writes the provided 'content' to a file with the specified 'filename' in the specified 'directory'.
*
*  If the specified 'directory' does not exist, it attempts to create it with read, write, and execute permissions for all users.
*  It then calls the 'writeFile' function to perform the file writing operation.
*
*  @param directory (string): The directory where the file will be created.
*  @param filename (string): The name of the file to be created.
*  @param content ([]byte): The content to be written to the file.
*/func WriteSimpleFile(directory, filename string, content []byte) {
	if !FileExists(directory) {
		utils.HandleBasicError(os.MkdirAll(directory, 0777), consts.ERR_CANT_CREATE_DIR)
	}
	createNewFile(content, directory+filename)
}
