# GO Starter

A GO starter project to provide a ready-to-implement codebase.

---

## Getting started

First, install the SDK: https://go.dev/doc/install  
Then, run the development server:

```bash
# Run the dev server
go run .
```

---

## CLI commands

Some of the CLI available commands:

- Start the GO project
    - `go run .`

- Create a new module
    - `go mod init <MODULE_NAME>`
    - e.g. `go mod init example.com/mymodule`

- Sync the project's dependencies analyzing the go.mod file
    - `go mod tidy`

- Build the project
    - `go build`

- Build and move an executable file in the $GOPATH/bin folder
    - `go install`

- Run the unit test
    - `go test`

- Show the help of the unit test
    - `go help test`

- Clean the packages cache
    - `go clean -modcache`
