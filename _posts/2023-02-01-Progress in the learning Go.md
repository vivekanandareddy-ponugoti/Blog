---
title: "Progress in learning GoLang"
date: 2023-02-1
---

# Progress in learning a new language:
I began learning Go by reading the official Go documentation, watching YouTube tutorials, and working on some small projects.
First, I had to install GoLang on my Mac from <a href="https://go.dev/dl">Download Go</a> website and installed it on my laptop.


## Learning basics of Go Language:
At first i created a .go file in visual studio code and implemented a Hello World!<a href="https://github.com/vivekanandareddy-ponugoti/Blog/blob/main/code/basics/helloWorld.go"></a>Hello World!</a> program in it.
I learned the command to run a go program: <code>go run fileName.go</code>

## Experimenting with packages in Go
After writing my first Hello World program, I wanted to learn how to work with multiple programs. Here, I noticed that programs in Golang are divided into packages just like in any other object oriented programming language. Each package can contain more than one function. The project is initialized in the main function, and we can use functions from other packages in the module by using the import command.
To initialize a module in Go we use command <code>go mod init module_name</code>
To import a module in Go we use command <code>import moduleName/packageName</code>

## Conclusion:
As I continue to learn Golang, I have an understanding of creating a new project or module using <code>go mod init</code>, The starting point of the project will be the main function of that module, which is located in <code>package main</code>.We can declare more than one function in a package, and it is very easy to import or export these functions from one package to another.