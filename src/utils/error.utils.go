package utils

func HandleBasicError(err error, message string) {
    if err != nil {
        println(message)
        panic(err)
    }
}
