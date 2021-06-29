package main

import "fmt"

/*
One of the most commonly used interfaces in
the Go standard library is the fmt.Stringer interface:

type Stringer interface {
	String() string
}
*/


type Article struct {
	Title string
	Author string
}

// (a Article): receiver, String() must be called by a type Article
// String(): the function name, with no arguments
// string: the return type
// Implements the String() method
func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s", a.Title, a.Author)
}

func main() {
	a := Article{
		Title: "Understanding Interfaces in Go",
		Author: "Jacinda Zhong",
	}

	Print(a)
}

func Print(s fmt.Stringer) {
	fmt.Println(s.String())
}
