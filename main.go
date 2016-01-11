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

    //参考一下这个网址, 去掉重复文件的.http://golanghome.com/post/476
    //还有一个网址:http://golangtc.com/t/54571b25421aa960c7000027
    root := "E:/111" //"H:/winxp"
    err := filepath.Walk(root, visit)
    fmt.Printf("filepath.Walk() returned %v\n", err)

    fmt.Println("end")
}
