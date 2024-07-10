# Contributing to fpkit

感谢你对 `fpkit` 项目的关注！我们欢迎各种形式的贡献，包括代码、文档、bug报告、功能请求等。为了让你的贡献顺利被接受，请遵循以下指导原则。

## 目录

- [Contributing to fpkit](#contributing-to-fpkit)
  - [目录](#目录)
  - [报告问题](#报告问题)
  - [提交变更](#提交变更)
  - [代码风格](#代码风格)
  - [测试](#测试)
  - [文档](#文档)
  - [许可证](#许可证)

## 报告问题

如果你发现了bug或有功能请求，请按照以下步骤操作：

1. 确保问题尚未被报告。可以通过搜索[Issues](https://github.com/mizumoto-cn/fpkit/issues)页面来确认。
2. 提供详细的信息，包括复现步骤、期望行为和实际行为。如果可能，提供相关代码片段。

## 提交变更

我们使用GitHub进行版本控制，请遵循以下步骤来提交你的变更：

1. Fork本仓库并clone到本地。
2. 创建一个新的分支：

    ```sh
    git checkout -b my-feature-branch
    ```

3. 在你的分支上进行开发，并确保代码符合代码风格的要求。
4. 添加测试，并确保所有测试都通过。
5. 提交变更并推送到你的fork：

    ```sh
    git commit -am "Add new feature"
    git push origin my-feature-branch
    ```

6. 创建一个新的Pull Request, 并描述你的变更项目与详细信息。

## 代码风格

我们使用Go语言进行开发，因此请遵循Go语言的[代码风格](https://golang.org/doc/effective_go.html)与[Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)

在提交前，请确保代码通过`gofmt`格式化，并通过`golangci-lint run`检查。

```sh
gofmt -w .
golangci-lint run
```

## 测试

我们使用`go test`进行测试，确保你的变更通过所有测试。

```sh
go test -race ./...
```

如果你的变更包含了新的功能，请添加对应的测试用例。

## 文档

我们使用Markdown格式进行文档编写，确保你的文档符合Markdown的语法规范。

本项目使用[DavidAnson/markdownlint](https://github.com/DavidAnson/markdownlint)进行Markdown文档的检查。

## 许可证

通过贡献代码，你同意将你的代码授权给项目的许可证[Mizumoto.General.Public.Licence.v1.5](./LICENSE)并接受其条款。详细信息请参考[许可证原文](./licensing/Mizumoto.General.Public.License.v1.5.md)。

再次感谢你的贡献！
