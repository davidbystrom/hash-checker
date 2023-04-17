package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
  if len(os.Args) != 3 {
    fmt.Println("Usage: ./main.go <file path> <sha256 hash>")
    os.Exit(1)
  }
  f, err := os.Open(os.Args[1])
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()

  h := sha256.New()
  if _, err := io.Copy(h, f); err != nil {
    log.Fatal(err)
  }


  hash := hex.EncodeToString(h.Sum(nil))

  fmt.Println("Generated hash from file: ", hash)
  fmt.Println("Hash collected from input:", strings.ToLower(os.Args[2]))
  fmt.Println("Hashes match:", string(hash) == strings.ToLower(os.Args[2]))
}