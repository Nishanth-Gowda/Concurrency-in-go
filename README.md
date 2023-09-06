# Concurrency in Go

![Go Logo](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/320px-Go_Logo_Blue.svg.png)

Welcome to the "Concurrency in Go" repository! This repository explores the concept of Concurrency with a focus on their implementation in the Go programming language.

## Table of Contents
- Introduction to goroutines(#introduction-to-goroutines)



## Introduction to goroutines
- A **Process** is an execution environment that contains instructions, user data, and system
data parts, as well as other types of resources that are obtained during runtime.

- A **Thread** is a smaller and lighter entity than a process. Threads are created by
processes and have their own flow of control and stack.

- A **goroutine** is the minimum Go entity that can be executed concurrently. The main
advantage of goroutines is that they are extremely lightweight and running thousands or
hundreds of thousands of them on a single machine will pose any problem.

## The Go scheduler
- Go runtime has its own scheduler which is responsible for the execution of the goroutine.

- The technique is also known as m:n scheduling, where m i goroutines are executed using n os threads.

- The Go scheduler is the Go component responsible for the way and the order in which the
goroutines of a Go program get executed.

## Goroutines
- You can define, create, and execute a new goroutine using the go keyword followed by a
function name.

- The go keyword makes the function call return immediately, while the function starts running 
in the background as a goroutine and the rest of the program continues its execution.

## Channels
- A channel is a communication mechanism that allows goroutines to exchange data, among
other things.

- Firstly, each channel allows the exchange of a particular data type, which is also called the e
lement type of the channel, and secondly, for a channel to operate properly, you will need someone to 
receive what is sent via the channel. 

- You should declare a new channel using the chan keyword, and you can close a channel using the 
close() function.
