# compiler written in Go
learning compiler and Go

[参考链接](http://www.craftinginterpreters.com/contents.html)

[Go notes](#go)
- [types](#types)
- [variable](#variable)
- [control flow](#control_flow)
- [function](#function)
- [input](#input)
- [output](#output)
- [run go](#run_go)
- [error handling](#error_handling)
- [Q&A](#q&a)
- [fucking stupid error i made QAQ](#fucking_stupid_error_i_made_QAQ)
[compiler notes](#compiler)

## Go
### types
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

#### struct
```Go=
type Book struct{
    name string
    price int
}

func (b *Book) getName() string{
    return b.name
}

```

#### enum
```Go=
const(
	A int = iota
	B
	C
	D
	)
```


#### interface
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
### variable
```Go=
// 1 init without init value
var a int
// 2 type is determined by literal
var a = 1
// 3
a := 1
```
### control flow
#### while does not exist
#### for
```Go=
for ; ; {

}

for true {

}
```

### function
```Go=
func hello(name string) string{
    return "hello " + name
}
```

#### overload
does not exist

#### optional parameter
does not exist

### input
#### read file
```Go=
str, err := ioutil.ReadFile(filename)
```

#### read stdin
```Go=
var input string
fmt.Scanln(&input)
```

#### read command line arguments
```Go=
arg := os.Args
```

### output
#### console output
```Go=
fmt.Println("hello world")
```

### run go
```
go build main.go
go run main.go
```
### error handling
no exception
- multiple returns
- panic
### Q&A
#### Java interface vs. Golang interface ?
[参考链接](https://stackoverflow.com/questions/39932713/whats-the-differences-between-go-and-java-about-interface)

### fucking stupid error i made QAQ
#### use pointer type when define method on struct!!! otherwise it will not change field
```
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


## compiler
### parse
#### solve ambiguous
precedence and associativity

#### tech
- recursive decent top-down
terminal match 上了就不会再回退了 之前还以可能还会回退 就有可能需要把token再吐出来23333
- LR bottom-up