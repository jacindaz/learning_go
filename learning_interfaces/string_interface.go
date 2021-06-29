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

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s", a.Title, a.Author)
}

func main() {
	a := Article{
		Title: "Understanding Interfaces in Go",
		Author: "Jacinda Zhong",
	}

	fmt.Println(a.String())
}
