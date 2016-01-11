package main
import (
    "os"
    "fmt"
    "path/filepath"
)

func visit(path string, fi os.FileInfo, err error) error {
    fmt.Printf("Visited: %s\n", path)
    if fi.Mode().IsRegular() {
        fmt.Println(fi.Name(), fi.Size(), "bytes")
    }
    return nil
}

func main(){

    root := "E:/111" //"H:/winxp"
    err := filepath.Walk(root, visit)
    fmt.Printf("filepath.Walk() returned %v\n", err)

    fmt.Println("end")
}
