package main

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark"
)

// This calls a JS function from Go.
func main() {}

// ...omitted

// This function is exported to JavaScript, so can be called using
// exports.multiply() in JavaScript.
//
//export Modifier
func Modifier(content string) string {
	md := []byte(content)
	var buf bytes.Buffer
	if err := goldmark.Convert(md, &buf); err != nil {
		panic(err)
	}
	println(content)
	fmt.Println("golang")
	return "from golang!!"

}
