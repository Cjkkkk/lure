# compiler written in Go
learning compiler and Go

[参考链接](http://www.craftinginterpreters.com/contents.html)

## 编译
根目录下

`go build -o lure.exe`

## 运行
### REPL
`lure.exe`

### 执行脚本
`lure [filename]`

## [Go notes](#go)
- [types](#types)
- [variable](#variable)
- [control flow](#control-flow)
- [function](#function)
- [input](#input)
- [output](#output)
- [run go](#run-go)
- [error handling](#error-handling)
- [Q&A](#qa)
- [fucking stupid things i wish i knew!](#fucking-stupid-things-i-wish-i-knew)

[compiler notes](#compiler)

## Go
### Types
- bool
- numeric
    - uint8
    - uint16
    - uint32
    - uint64
    - int8
    - int16
    - int32
    - int64
    - float32
    - float64
    - complex64
    - complex128
- string
- derived
    - Pointer types
    - Array types
    - Structure types
    - Union types and
    - Function types
    - Slice types
    - Interface types
    - Map types
    - Channel Types

- #### struct
```Go=
type Book struct{
    name string
    price int
}

func (b *Book) getName() string{
    return b.name
}

```

- #### enum
```Go=
const(
	A int = iota
	B
	C
	D
	)
```


- #### interface
```
type geometry interface{
    area() float64
}

type circle struct{
    radius float64
}

func (c circle) area() float64{
    return 3.14 * c.radius * c.radius
}
```
### Variable
```Go=
// 1 init without init value
var a int
// 2 type is determined by literal
var a = 1
// 3
a := 1
```
### control flow
- #### while does not exist
- #### for
```Go=
for ; ; {

}

for true {

}
```

### Function
```Go=
func hello(name string) string{
    return "hello " + name
}
```

- #### overload
does not exist

- #### optional parameter
does not exist

### Input
- #### read file
```Go=
str, err := ioutil.ReadFile(filename)
```

- #### read stdin
```Go=
var input string
fmt.Scanln(&input)
```

- #### read command line arguments
```Go=
arg := os.Args
```

### Output
- #### console output
```Go=
fmt.Println("hello world")
```

### Run go
```
go build main.go
go run main.go
```
### error handling
no exception
- multiple returns
- panic

- #### Custom error
```
type MyError struct{

}
func (e *MyError) Error() string {
    return "I am an error"
}
```
- #### basic error
```
errors.New("can't work with 42")
```
### Q&A
- #### Java interface vs. Golang interface ?
[参考链接](https://stackoverflow.com/questions/39932713/whats-the-differences-between-go-and-java-about-interface)

- #### how to handle error in Go?

### Fucking stupid things i wish i knew!
- #### Use pointer type when define method on struct!!! otherwise it will not change field
```Go=
type animal struct{
    age int
}
func (a animal) aging() {
    a.age ++
}

func (a *animal) aging_p() {
    a.age ++
}

var a = animal{age: 1}
a.aging()
fmt.Println(a.age) // 1
a.aging_p()
fmt.Println(a.age) // 2
```
- #### Always use uppercase in package
always use uppercase in package if you want to export that to other package

## compiler
### parse
#### solve ambiguous
- precedence and associativity
- add more non-terminal to introduce precedence  
#### tech
- recursive decent top-down
terminal match 上了就不会再回退了 之前还以可能还会回退 就有可能需要把token再吐出来23333
- LR bottom-up

#### runtime error
- when convert tokens into AST, we throw static error
- when walking AST, we throw runtime error


#### three-address code
It is possible that a compiler will construct a syntax tree at the same time
it emits steps of three-address code. However, it is common for compilers to
emit the three-address code while the parser "goes through the motions" of
constructing a syntax tree, without actually constructing the complete tree
data structure

#### L-Value R-Value
Right now, the only valid target is a simple variable expression