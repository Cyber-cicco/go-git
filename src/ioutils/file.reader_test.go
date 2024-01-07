package ioutils

import (
    "testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {

    parentDir1 := walkToParentDirectory("/caca/pipi")
    parentDir2 := walkToParentDirectory("/caca/pipi\\/pipi")
    parentDir3 := walkToParentDirectory("/caca")

    if parentDir1 != "/caca" {
        t.Fatalf("c'est pas caca %q", parentDir1)
    }

    if parentDir2 != "/caca" {
        t.Fatalf("c'est pas caca %q", parentDir2)
    }

    if parentDir3 != "" {
        t.Fatalf("c'est pas caca %q", parentDir2)
    }
}
