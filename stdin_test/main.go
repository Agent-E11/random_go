package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        input, err := reader.ReadString('\n')
        if err != nil {
            if err == io.EOF {
                fmt.Println("End of file. Quitting...")
            } else {
                fmt.Printf("Error: %v", err)
            }
            return
        }
        if strings.TrimSpace(input) == "q" {
            fmt.Println("Quitting...")
            return
        }
        fmt.Printf("Your input: %s", input)
    }
}

