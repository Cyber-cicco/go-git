package utils

func IndexOf(content []byte, char byte) int {
    i := 0
    for i < len(content) {
        if content[i] == char {
            return i
        }
        i += 1
    }
    return -1
}
