# Contributing to fpkit

[中文版](./CONTRIBUTING_zh.md)

Thank you for your interest in contributing to the `fpkit` project! We welcome all forms of contributions, including code, documentation, bug reports, and feature requests. To ensure that your contributions are accepted smoothly, please follow the guidelines below.

## Table of Contents

- [Contributing to fpkit](#contributing-to-fpkit)
  - [Table of Contents](#table-of-contents)
  - [Reporting Issues](#reporting-issues)
  - [Submitting Changes](#submitting-changes)
  - [Code Style](#code-style)
  - [Testing](#testing)
  - [Documentation](#documentation)
  - [License](#license)

## Reporting Issues

If you find a bug or have a feature request, please follow these steps:

1. Ensure the issue has not already been reported. You can check this by searching the [Issues](https://github.com/mizumoto-cn/fpkit/issues) page.
2. Provide detailed information, including reproduction steps, expected behavior, and actual behavior. If possible, provide relevant code snippets.

## Submitting Changes

We use GitHub for version control. Please follow these steps to submit your changes:

1. Fork this repository and clone it to your local machine.
2. Create a new branch:

    ```sh
    git checkout -b my-feature-branch
    ```

3. Develop on your branch and ensure your code meets the code style guidelines.
4. Add tests and ensure all tests pass.
5. Commit your changes and push to your fork:

    ```sh
    git commit -am "Add new feature"
    git push origin my-feature-branch
    ```

6. Create a new Pull Request and describe your changes in detail.

## Code Style

We use Go for development, so please follow the Go language's [Effective Go](https://golang.org/doc/effective_go.html) and [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments).

Before submitting, ensure your code is formatted with `gofmt` and passes `golangci-lint` checks.

```sh
gofmt -w .
golangci-lint run
```

## Testing

We use `go test` for testing. Ensure your changes pass all tests.

```sh
go test -race ./...
```

If your changes include new features, please add corresponding test cases.

## Documentation

We use Markdown for documentation. Ensure your documentation follows Markdown syntax guidelines.

This project uses [DavidAnson/markdownlint](https://github.com/DavidAnson/markdownlint) for checking Markdown documents.

## License

By contributing code, you agree to license your code under the project's [Mizumoto General Public License v1.5](./LICENSE) and accept its terms. For more details, please refer to the [license text](./licensing/Mizumoto.General.Public.License.v1.5.md).

Thank you again for your contribution!
