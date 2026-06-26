# Learning Golang

In this repository, I will record my studies and learnings of the [Go](https://go.dev/) programming language.

## About

This repository contains examples, experiments, notes, and small projects developed while studying [Go](https://go.dev/).

The goal is to understand [Go](https://go.dev/) fundamentals, explore concepts, and build practical projects over time.

### Why I decided to learn Golang?

I decided to learn [Go](https://go.dev/) because of its performance, simplicity, and clean language design. I'm particularly interested in how this language is used to build backend applications and CLI tools. As someone who primarily works with TypeScript, I also wanted to explore a compiled language with a different approach to software development.

## Requirements

To build this project and use in your machine, make sure you have the following installed:

- **Go 1.26** - Thats the version that I've used
- **Git**

## Installation

### Debian/Ubuntu

```bash
sudo apt update
sudo apt install golang-go git
```

### Arch Linux

```bash
sudo pacman -S go git
```

### macOS

Using Homebrew:

```bash
brew install go git
```

### Windows

- Install Go: **https://go.dev/dl/**
- Install Git: **https://git-scm.com/downloads**

### Verify the installation:

```bash
go version
git --version
```

## Cloning this repository

To clone this repository, just run the following command:

```bash
git clone https://github.com/lucaaszsx/learning-go
cd learning-go
```

## Repository Structure

```
├── assets
├── build
├── fundamentals
│   ├── basics
│   │   ├── conditional
│   │   ├── constant
│   │   ├── error
│   │   ├── function
│   │   ├── hello
│   │   ├── input
│   │   ├── interface
│   │   ├── loop
│   │   ├── pointer
│   │   ├── struct
│   │   ├── types
│   │   └── variable
│   └── stdlib
│       ├── context
│       ├── encoding-json
│       ├── errors
│       ├── fmt
│       ├── io
│       ├── math
│       ├── net-http
│       ├── os
│       ├── regexp
│       ├── sort
│       ├── strconv
│       ├── strings
│       ├── sync
│       ├── testing
│       └── time
├── go.mod
├── LICENSE
├── projects
│   └── todo-api
├── README.md
└── scripts
```

## Running examples

After cloning this repository, you will be able to run the code examples.

For run the example `fundamentals/basics/types`, execute the following command:

```bash
go run fundamentals/basics/types/main.go
```

For run a stdlib example like `fundamentals/stdlib/fmt`:

```bash
go run fundamentals/stdlib/fmt/main.go
```

### Running projects

Projects have their own `go.mod`, so you need to enter the project directory first:

```bash
# Example: running the todo-api project
cd projects/todo-api
go run main.go
```

Each project has its own README with specific instructions, endpoints, and usage examples.

## Notes

This repository is intended for educational purposes and may contain code experiments, refactorings, and incomplete implementations.

## License

This repository is licensed under **MIT license**. To see the full license text, [go to LICENSE file](./LICENSE).
