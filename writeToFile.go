file, _ := os.Create("output.txt")
    fmt.Fprint(file, "This is how you write to a file, by the way")
    file.Close()
