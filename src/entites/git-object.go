package entites

import (
	"fmt"
	"fr/eduprolo/src/ioutils"
)

type GitObject interface {
    Serialize(file string) string
    Deserialize() string
}

type GitFlag struct {
    Length int
    Content string
}

type GitCommit struct {
    Tree string
    Parent string
    Author string
    Commiter string
    GpgSig string
}

type GitTree struct {
    Length int
    Content string
}
type GitBlob struct {
    Length int
    Content string
}

func (g *GitFlag) Serialize(file string) string{
    return ""
}

func (g *GitFlag) Deserialize() string {
    return g.Content
}
func (g *GitCommit) Serialize(file string) string {
    return ""
}

func (g *GitCommit) Deserialize() string {
    return ""
}
func (g *GitTree) Serialize(file string) string {
    return ""
}

func (g *GitTree) Deserialize() string {
    return g.Content
}
func (g *GitBlob) Serialize(file string) string {
    fileContent := ioutils.ReadFile(file)
    fileLen := len(fileContent)
    gitContent := fmt.Sprintf("blob %d%c%s", fileLen, '\x00', fileContent)
    return gitContent
}

func (g *GitBlob) Deserialize() string {
    return g.Content
}
