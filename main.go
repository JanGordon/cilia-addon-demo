package main

import "github.com/gomarkdown/markdown"

// This calls a JS function from Go.
func main() {}

// ...omitted

// This function is exported to JavaScript, so can be called using
// exports.multiply() in JavaScript.
//
//export Modifier
func Modifier(content string) string {
	md := []byte(content)
	output := markdown.ToHTML(md, nil, nil)
	return string(output)

}
