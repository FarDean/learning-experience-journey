## Go code organization

**package**: is a collection of source files in the same directory.

**module**: is a collection of Go packages that are released together.
a Go repo typically contains only one module.

**go.mod**: a file that declares the `module path`, the import path for all packages within the module.


## First step

1. create a `go.mod` file:
```sh
$ go mod init [path]/hello
```
- The first statement in a Go source file must be package name. Executable commands must always use package main.

1. create a file named `hello.go`.
2. build and install the program with:
```sh
$  go install [path]
```
- This command builds the `hello` command, producing an executable binary. It then installs that binary as $HOME/go/bin/hello (or, under Windows, %USERPROFILE%\go\bin\hello.exe).
- The install directory is controlled by the `GOPATH` and `GOBIN` environment variables. If `GOBIN` is set, binaries are installed to that directory. If `GOPATH` is set, binaries are installed to the `bin` subdirectory of the first directory in the `GOPATH` list. Otherwise, binaries are installed to the `bin` subdirectory of the default `GOPATH` ($HOME/go or %USERPROFILE%\go).
- `go env` command can set the default environment variable:
```sh
$ go env -w GOBIN=/somewhere/else/bin
```
- To unset a variable previously set by go env -w, use go env -u:
```sh
$ go env -u GOBIN
```
- `go` commands accept paths relative to the working directory,So in our working directory, the following commands are all equivalent:
```sh
$ go install [path]
```
```sh
$ go install .
```
```sh
$ go install 
```
3. run the command to make sure the module is working:
```sh
$ hello
```

[Refrence](https://go.dev/doc/code)