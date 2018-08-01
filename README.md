# iGo - a Golang interpreter, written in Go

![iGo-Demo](https://static.beeceej.com/iGoDemo.gif)

**VSCode Integration still in pre-alpha Development**

## Usage

`$ cd $path_to_iGo`
`$ go run cmd/interpreter/main.go` This will start up the interpreter server

The interpreter will then parse your go code. And if possible evaluate the expression
`$ curl -XPOST http://localhost:9999/interpret -d'{"text": $your_go_code}'`

## Supported

- Function Parsing
  - Single Function parsing
  - Multiple Function parsing at a time
- Expressions
  - Calling built in functions
  - Calling user defined functions
