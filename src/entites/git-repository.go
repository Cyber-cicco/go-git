package entites

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
