package main

import (
    "fmt"
    "github.com/prasad-firame/mirro/internal/proxy"
)

func main() {
    fmt.Println("Starting Mirro reverse proxy...")
    proxy.Start(":8080")
}