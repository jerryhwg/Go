# Go
Learning Go

# Install

https://golang.org/dl

In VS Code

Search Go in Extensions and install

> Right bottom -> Plain Text -> Go -> "Some Go analysis tools are missing from your GOPATH. Install.

Env setup

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
export PATH="/usr/local/go/bin:$PATH"
```

Check

```bash
go version
```

## Run

```bash
go run main.go
```

## Commands

```bash
go build
# compiles a bunch of go source code files

go run
# compiles and executes one or two files (no build)

go fmt
# formats all the code in each file in the current directory

go install
# compiles and "installs" a package
# build the app and put the app library in the $GOPATH/bin directory

go generate
# generate or update documentation

go get
# downloads the raw source code of someone else's package

go test
go test -v ./...
# runs any tests associated with the current project
```

Go build

```bash
go build
# run
./main
```

## Basics

package = project = workspace

* executable: package main
* re-usable: package < name > : dependency, library

array: fixed

slice: grow

## Types

### Value Types

> int, float, string, bool, structs

Use pointers to change these things in a function

### Reference Types

> slices, maps, channels, pointers, functions

Don't worry about pointers with these

## Struct vs. Map

### Struct

* (Similar to dictionary in Python)
* Values can be of different type
* Keys don't support indexing
* Use to represent a "thing" with a lot of different properties
* You need to know all the different fields at compile time
* Value Type!

### Map

* Collection of key value pair
* (Similar to dictionary in Python)
* All keys must be the same type
* All values must be the same type
* Keys are indexed - we can iterate over them
* Use to represent a collection of related properties
* Don't need to know all the keys at compile time
* Reference Type!
