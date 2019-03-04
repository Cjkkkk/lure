# compiler written in Go
learning compiler and Go

[参考链接](http://www.craftinginterpreters.com/contents.html)

[Go notes](#go)

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
### variable initialization
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


## compiler