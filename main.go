package main

import (
	"fmt"
	"strings"
)

func main() {
    fmt.Println("Hello!")
    s := []string{
        "This", "is", "a", "test", "of",
        "string", "filtering", "and", "interfaces",
    }

    filtered := filter(s, func(e string) bool { return strings.Contains(e, "i")})
    fmt.Printf("%v\n", filtered)
}

func filter[T any](slice []T, test func(T) bool) []T {
    ret := []T{}

    for _, element := range slice {
        if test(element) {
            ret = append(ret, element)
        }
    }
    return ret
}
