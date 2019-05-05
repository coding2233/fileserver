# fileserver
简单的文件服务器

一句代码实现文件服务器
```go 
// Go语言 net/http
http.ListenAndServe(":"+port, http.FileServer(http.Dir("./files")))
```


### 使用步骤

#### windows x64
* 运行`fileserver.exe`,默认端口`8091`
* 写一个`run.bat`文件,可动态设置端口
```bat
@echo off
fileserver.exe 8092
```

#### linux x64
* 直接运行`./fileserver`，默认端口`8091`
* 动态设置端口`./fileserver 8092`
