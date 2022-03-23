# study-go
关于一些go语言学习
## 常用命令
- go version
- go help
- go mod init
- go mod tidy
- go clean -modcache
- go run .
- go build
- go install
- go test -v
- go test -fuzz=Fuzz -fuzztime 30s
- go get .
- go doc -all regexp | grep -i parse
- go fmt
## Effective go
- 命名 驼峰式，可被外部包访问的函数名称首字母大写
- 分号 词法分析器会自动插入，不应将一个控制结构（if、for、switch 或 select）的左大括号放在下一行
- defer 延迟执行，先进后出，用于释放资源/关闭文件等
## 语法
  - 声明变量 `var`,可以一次声明多个变量`var a int = 1`, `var c, d = "hello", "world"`；`:=`也是声明语句，`e:=3`,出现在`:=`左侧的变量不应该是已经被声明过的，否则会导致编译错误。全局变量允许声明但不使用，但局部变量会报错。
  - 值类型和引用类型 值类型的变量直接指向存在内存中的值，`j = i`实际上是在内存中将 i 的值进行了拷贝。引用类型的变量存储的是变量的值所在的内存地址（数字），或内存地址中第一个字所在的位置。这个内存地址称之为指针，这个指针实际上也被存在另外的某一个值中。 `r2 = r1`只有引用（地址）被复制。
