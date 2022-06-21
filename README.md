# study-go

关于一些golang的学习笔记:

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

- 声明变量 `var`,可以一次声明多个变量`var a int = 1`, `var c, d = "hello", "world"`；`:=`也是声明语句，`e:=3`,出现在`:=`
  左侧的变量不应该是已经被声明过的，否则会导致编译错误。全局变量允许声明但不使用，但局部变量会报错。
- 值类型和引用类型 值类型的变量直接指向存在内存中的值，`j = i`实际上是在内存中将 i
  的值进行了拷贝。引用类型的变量存储的是变量的值所在的内存地址（数字），或内存地址中第一个字所在的位置。这个内存地址称之为指针，这个指针实际上也被存在另外的某一个值中。 `r2 = r1`只有引用（地址）被复制。
- 运算符 `<<`乘以2的n次方，`>>`除以2的n次方。
- 切片是对数组的抽象，数组的长度不可变，与数组相比切片的长度是不固定的("动态数组")，可以追加元素，在追加时可能使切片的容量增大。
- 类型转换 不支持隐式转换，go是强类型语言。
- 并发 goroutine 是轻量级线程，同一个程序中的所有 goroutine 共享同一个地址空间；通道（channel）是用来传递数据的一个数据结构。


## 框架

1. 项目目录
   - /pkg：可以被外部应用使用的代码库
   - /internal：私有代码
     - /app：项目代码
     - /pkg：内部应用依赖的代码库，如pg、redis
   - /src：/model、/ctr都不允许存在
   - /cmd：可执行文件
   - /api：api接口定义
   - Makefile：触发脚本
   - scripts：需要运行的脚本
2. 按照职责对模块进行垂直拆分，容易对微服务进行拆分。而不是按层级水平拆分。
3. 显式与隐式
  - init 显式。在包被引入时隐式地执行一些代码，慎重使用。只做一些简单轻量的前置条件判断，不做过重的初始化逻辑。
  - error 显式。
4. 面向接口：单元测试是一个项目保证工程质量最有效并且投资回报率最高的方法之一
   - 使用大写的 Service 对外暴露方法；
   - 使用小写的 service 实现接口中定义的方法；
   - 通过 func NewService(...) (Service, error) 函数初始化 Service 接口；
5. 单元测试：单元测试就是业务逻辑。
    - Test 

    单元测试的执行不应该依赖任何外部模块，单元测试不是集成测试，它的运行不应该依赖除项目代码外的其他任何系统。
    ```
    // add_test.go
    func TestAdd(t *testing.T) {
        tests := []struct{
            name     string
            first    int64
            second   int64
            expected int64
        } {
            {
                name:     "HappyPath":
                first:    2,
                second:   3,
                expected: 5,
            },
            {
                name:     "NegativeNumber":
                first:    -1,
                second:   -1,
                expected: -2,
            },
        }
        
        for _, test := range tests {
            t.Run(test.name, func(t *testing.T) {
                assert.Equal(t, test.expected, Add(test.first, test.second))
            })
        }
    }
    ```

    - Mock

        - gomock
   
       `$ mockgen -package=mblog -source=pkg/blog/blog.go > test/mocks/blog/blog.go`
  
      - sqlmock
      - httpmock
    - 断言