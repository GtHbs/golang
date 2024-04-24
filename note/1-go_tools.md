**go build**
> 编译工具，可以将包和相关的依赖编译成一个可执行性文件
> 
> `go build [-o output] [-i] [build flags] [packages]`
> 
> go build 需要的参数是一个路径，packages是一个相对路径，是相对于`GOROOT`和`GOPATH`两个环境变量

> 

**go clean**
> 用于清理生成的可执行性文件
> 
> 参数：
> 
> `go clean [-i] [-r] [-n] [-x] [build flags] [packages]`
> 
> -i：清理安装在指定packages下的可执行性文件


**go run**
> 将`go build`和可执行性文件执行合二为一，但必须跟随go文件参数
> 
> 参数：
> 
> `go run *.go`

**go env**
> go环境
> 
> GOOS：指编译后要适应的操作系统，可选值
> darwin,freebsd,linux,windows,android,dragonfly,netbsd,openbsd,plan9,solaris
> 
> GOARCH：目标处理器架构，可选值
> arm,arm64,386,amd64,ppc64,ppc64le,mips64,mips64le,s390x


**go install**
> 打包、编译文件后，将程序安装到固定目录
> 
> 参数：`go install [build flags] [packages]`


**go get**
> 用于下载相关包，并且安装到本地

**go vet**
> 用于检查代码中常见的错误
> 
> 参数：`go vet [-n] [-x] [build flags] [packages]`

**go test**
> 用于代码的单元测试，接收一个包作为参数。
> 
> 要求
> - 代码文件必须以`_test.go`结尾。
> - 测试代码必须包含若干个测试函数。
> - 测试函数必须以`Test`开头，接收`*testing.T`的参数