package main

import (
	"io"
	"os"
	"testing"
)

func TestPushAndPop(t *testing.T) {
    stack := NewStack()
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)

    popped := stack.Pop()
    if popped != 3 {
        t.Errorf("Expected popped value to be 3, but got %d", popped)
    }

    popped = stack.Pop()
    if popped != 2 {
        t.Errorf("Expected popped value to be 2, but got %d", popped)
    }

    popped = stack.Pop()
    if popped != 1 {
        t.Errorf("Expected popped value to be 1, but got %d", popped)
    }

    // Test popping from an empty stack
    popped = stack.Pop()
    if popped != -1 {
        t.Errorf("Expected popped value to be -1 for an empty stack, but got %d", popped)
    }
}

func TestPrint(t *testing.T) {
    stack := NewStack()
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)

    // Redirect stdout to capture printed output
    capturedOutput := captureOutput(func() {
        stack.Print()
    })

    expectedOutput := "1\n2\n3\n"
    if capturedOutput != expectedOutput {
        t.Errorf("Expected printed output:\n%s\nBut got:\n%s", expectedOutput, capturedOutput)
    }
}

func captureOutput(f func()) string {
    // Capture stdout
    originalOutput := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    // Execute the function that prints
    f()

    // Restore stdout
    w.Close()
    os.Stdout = originalOutput

    // Read captured output
    capturedOutput, _ := io.ReadAll(r)

    return string(capturedOutput)
}
