package entites

type GitObject interface {
    Serialize() 
    Deserialize() string
}

type GitFlag struct {
    Length int
    Content string
}
type GitCommit struct {
    Length int
    Content string
}
type GitTree struct {
    Length int
    Content string
}
type GitBlob struct {
    Length int
    Content string
}

func (g *GitFlag) Serialize() {
}

func (g *GitFlag) Deserialize() string {
    return g.Content
}
func (g *GitCommit) Serialize() {
}

func (g *GitCommit) Deserialize() string {
    return g.Content
}
func (g *GitTree) Serialize() {
}

func (g *GitTree) Deserialize() string {
    return g.Content
}
func (g *GitBlob) Serialize() {
}

func (g *GitBlob) Deserialize() string {
    return g.Content
}
