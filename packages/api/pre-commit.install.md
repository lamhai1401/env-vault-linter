# Overview

Welcome to Pre-commit.
It's nice feature for linter, clean code, testing, better comment for project you developer.

## Install

Please visit to [Pre-commit](https://pre-commit.com), [golangci-lint](https://golangci-lint.run/usage/install/) to choose which version is suitable for your computer.


## How to use it

Go to root folder of your project and run these command

```bash
pre-commit install
```

## How to work with it

***The first run will take time***

It just like a before git commit hook.

Every time you commit a new feature, it will check all of it and prevent commit if it fail the test (linter, grammar, func never use) and sometime will suggest you how to fix it.

You have to fix all of issue it found.
Happy coding :)

## Install go pkg to run

```go
go install github.com/mgechev/revive@master
go install github.com/securego/gosec/v2/cmd/gosec@latest
go install github.com/orijtech/structslop/cmd/structslop@latest
go install mvdan.cc/gofumpt@latest
go install golang.org/x/tools/cmd/goimports@latest
go install -v github.com/go-critic/go-critic/cmd/gocritic@latest
````
