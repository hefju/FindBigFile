package main
import (
    "os"
    "fmt"
    "errors"
    "path/filepath"
)

func visit(path string, fi os.FileInfo, err error) error {

    if fi==nil{
        //遇到的错误: c:/hiberfil.sys: The process cannot access the file because it is being used by another process.
        return errors.New("bad file: "+path)
    }
    if fi.Size()<gb {//panic: runtime error: invalid memory address or nil pointer dereference
        return nil
    }

    fmt.Printf("Visited: %s\n", path)
    if fi.Mode().IsRegular() {
        fmt.Println(fi.Name(), fi.Size(), "bytes")
    }
    return nil
}
 var gb int64 =1000*1000*1000*0.5
func main(){

    //参考一下这个网址, 去掉重复文件的.http://golanghome.com/post/476
    //还有一个网址:http://golangtc.com/t/54571b25421aa960c7000027
    root := "c:/" //"H:/winxp"
    err := filepath.Walk(root, visit)
    fmt.Printf("filepath.Walk() returned %v\n", err)

    fmt.Println("end")
}
