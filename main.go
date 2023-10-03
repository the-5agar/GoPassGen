// main.go
package main

import (
    "fmt"
    "gopassgen/cmd"
)

func main() {
    if err := cmd.Execute(); err != nil {
        fmt.Println(err)
    }
}
