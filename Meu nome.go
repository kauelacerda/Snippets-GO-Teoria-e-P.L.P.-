package main

import "fmt"

type Speaker interface {
	Speak() string
}

type User struct {
	Name string
}

// User implicitly satisfies Speaker without an "implements" keyword
func (u User) Speak() string {
	return "Hello, my name is " + u.Name
}

func Announce(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	u := User{Name: "Alex"}
	Announce(u)
}
