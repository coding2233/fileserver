package main

import (
    "net/http"
    "fmt"
    "os"
)

func main() {

    args :=  os.Args

    //默认端口8091
    port := "8091"

    //从外面取参数 做端口
    if len(args) > 1{
        port=args[1]
    }

    // http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("./tmp"))))

    // err := http.ListenAndServe(":8090", nil)
    // if err != nil {
    //     fmt.Println(err)
    // }
    
    fmt.Println(":"+port)
    fmt.Println("./files ")
    fmt.Println("\nhttp file server runing... ")

    http.ListenAndServe(":"+port, http.FileServer(http.Dir("./files")))
}