package utils

import "github.com/elliotchance/orderedmap/v2"

func ParseKeyValueFile(fileContent string) *orderedmap.OrderedMap[string, string] {
    keyValMap := orderedmap.NewOrderedMap[string, string]()
    parseKeyValueFile(fileContent, 0, keyValMap)
    return keyValMap
}

func parseKeyValueFile(fileContent string, pos int, keyValMap *orderedmap.OrderedMap[string, string]) {
    idx := pos
    hasWhiteSpace := false
    posOfWhiteSpace := pos
    startPos := pos
    for idx < len(fileContent) {
        if fileContent[idx] == ' ' && idx == pos {
            for fileContent[idx] != '\n' && idx < len(fileContent) {
                idx++
            }
            idx++
            pos = idx
            continue
        }

        if fileContent[idx] == ' ' {
            hasWhiteSpace = true
            posOfWhiteSpace = idx
            idx++
            continue
        }
        if fileContent[idx] == '\n' && hasWhiteSpace {
            key := fileContent[startPos:posOfWhiteSpace]
            val := fileContent[posOfWhiteSpace+1:idx]
            keyValMap.Set(key, val)
            idx++
            parseKeyValueFile(fileContent, idx, keyValMap)
        }
        if fileContent[idx] == '\n' {
            idx++
            parseKeyValueFile(fileContent, idx, keyValMap)
        }
        idx++
    }
}
