package entites

type GitObject interface {
    Serialize(repo *GitRepository) 
    Deserialize(repo *GitRepository) 
}

type GitBlob struct {
    Length int
    Content string
}

func (g *GitBlob) Serialize() {
}

func (g *GitBlob) Deserialize() {
}
