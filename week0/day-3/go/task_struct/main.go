package main

import "fmt"

type Task struct {
	Title    string
	Done     bool
	Priority int
}

func (t *Task) Complete() {
	// TODO
}

func (t *Task) Bump() {
	// TODO
}

func (t Task) String() string {
	// TODO
	return ""
}

func main() {
	t := Task{Title: "Learn Go", Priority: 1}
	fmt.Println(t)
	t.Bump()
	t.Complete()
	fmt.Println(t)
}
