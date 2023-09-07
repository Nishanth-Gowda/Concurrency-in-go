package main

import "fmt"

type Stack interface {
    Push(item int)
    Pop() int
    Print()
}

type MyStack struct {
    data []int
}

func (s *MyStack) Push(item int) {
    s.data = append(s.data, item)
}

func (s *MyStack) Pop() int {
    l := len(s.data)
    if l == 0 {
        return -1
    }
    item := s.data[l-1]
    s.data = s.data[:l-1]
    return item
}

func (s *MyStack) Print() {
    for _, v := range s.data {
        fmt.Println(v)
    }
}

// NewStack creates and returns a new instance of MyStack.
func NewStack() *MyStack {
    return &MyStack{}
}

func main() {
    stack := NewStack()
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)

    stack.Print()

    fmt.Println("Popped item:", stack.Pop())

    stack.Print()
}
