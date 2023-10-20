package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "path/filepath"
    "math/rand"
    "time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const randomStringLen = 16

func main() {
    rand.Seed(time.Now().UnixNano())
    files, err := ioutil.ReadDir(".")
    if err != nil {
        return nil, fmt.Errorf("Error reading the current directory: %w", err)
    }

    for _, file := range files {
        if file.IsDir() {
            continue
        }

        // Check if the file has a supported image extension.
        ext := filepath.Ext(file.Name())
        if isImageExtension(ext) {
            newName := generateRandomString() + ext
            err := os.Rename(file.Name(), newName)
            if err != nil {
                return nil, fmt.Errorf("Error renaming file: %w", err)
            } else {
                fmt.Printf("Renamed %s to %s\n", file.Name(), newName)
            }
        }
    }
}

func isImageExtension(ext string) bool {
    // Add more image extensions if needed.
    imageExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff"}
    for _, imgExt := range imageExtensions {
        if ext == imgExt {
            return true
        }
    }
    return false
}

func generateRandomString() string {
    b := make([]byte, randomStringLen)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}
