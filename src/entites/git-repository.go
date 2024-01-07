package entites

import (
	"bytes"
	"encoding/json"
	"fr/eduprolo/src/consts"
	"fr/eduprolo/src/utils"
	"os"
)

type GitRepository struct {
    Worktree string
    GitDir string
    Conf ConfigurationFile
}

type ConfigurationFile struct {
    Core Core `json:"core"`
}

type Core struct {
    RepositoryFormatVersion int32 `json:"repositoryformatversion"`
    FileMode bool `json:"filemode"`
    Bare bool `json:"base"`
}

func GetBaseConfigFile() ConfigurationFile {
    return ConfigurationFile{
        Core: Core{
            RepositoryFormatVersion : 0,
            FileMode : false,
            Bare : false,
        },
    }
}

func GetGitRepoFromPath(path string) GitRepository {
    conf := unMarshalFile(path + "/.git/config")
    return GitRepository{
        Worktree : path,
        GitDir : path + "/.git/",
        Conf : conf,
    }
}

func unMarshalFile(path string) ConfigurationFile {
    f, err := os.ReadFile(path)
    utils.HandleBasicError(err, consts.ERR_CANT_READ_FILE)
    var demarshaledFile ConfigurationFile
    json.Unmarshal(f, &demarshaledFile)
    return demarshaledFile
}

func (g *GitRepository) GetObjectFileFromHash(sha string) bytes.Buffer {
    path := g.GitDir + consts.DIR_OBJECTS + sha[0:2] + "/" + sha[2:]
    content, err := os.ReadFile(path)
    utils.HandleBasicError(err, consts.ERR_CANT_READ_FILE)
    return *bytes.NewBuffer(content)
}

