This repository shows how to use [custom linters](https://github.com/BelehovEgor/golang-linter) for your project.

### Steps
1.  First of all you need to add package with linters and package with analysis tools in your `go.mod` file. Like this
    ```
    github.com/BelehovEgor/golang-linter <some version>
    golang.org/x/tools <some version>
    ```

2.  Then add cmd/analyzer/analyzer.go file in your project. You can see [example](./cmd/analyzer/analyzer.go) in this repository.
3.  In analyzer.go import necessary analyzers, configure and add them into `multichecker`
4.  Build analyzer with command `go build ./cmd/analyzers`
5.  Use resulted analyzer on your files. For example `./analyzers example.go` print this message
    ```
    /.../go-lint-integration/example.go:1:1: Don't use mutable global variable '_someData', use dependency injection.
    ```
