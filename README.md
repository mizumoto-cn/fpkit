# fpkit

![Go Version](https://img.shields.io/badge/Go-1.22.4-yellow.svg)
[![License](https://img.shields.io/badge/License-MGPL%20v1.5-green.svg)](/Licensing/Mizumoto.General.Public.License.v1.5.md)

[![Go Report Card](https://goreportcard.com/badge/github.com/mizumoto-cn/fpkit)](https://goreportcard.com/report/github.com/mizumoto-cn/fpkit)
[![CodeFactor](https://www.codefactor.io/repository/github/mizumoto-cn/fpkit/badge)](https://www.codefactor.io/repository/github/mizumoto-cn/fpkit)
[![codecov](https://codecov.io/github/mizumoto-cn/fpkit/graph/badge.svg?token=KB82WFBUQ3)](https://codecov.io/github/mizumoto-cn/fpkit)
[![Build](https://github.com/mizumoto-cn/fpkit/actions/workflows/go.yml/badge.svg)](https://github.com/mizumoto-cn//fpkit/actions)

[中文](./README_zh.md)

A light-weight Go functional tools lib, with generic programming support.

## Quick Start

First, install the package.

```bash
go get github.com/mizumoto-cn/fpkit@latest
```

This library is developed using Go 1.22.4 and later, and your Go version should be 1.18 or later.

We welcome your feedback and contributions.

## Documentation

All documentation is available in the [Wiki](./Wiki/) folder.

## Milestones

## Roadmap

### v0.0.1

- [x] Basic Slice: CRUD, Union, Intersection
- [x] Basic Functional Programming: Map, Filter, Reduce/Fold, Numeric, ...

### v0.0.2

- [x] Functional Programming: Optional/Maybe
- [x] Functional Programming: Currying, Composition, ...
- [ ] Basic Queue: Queue, Priority Queue

### v0.1.0

- [ ] Change Basic Slice to more functional programming style
- [ ] Basic Map
- [ ] Advanced Map: Hash map, Tree map, Linked map
- [ ] Set: Hash set, Tree set, Sorted set
- [ ] Advanced Queue and Concurrency: Concurrent queue, Concurrent blocking queue, Concurrent blocking priority queue

### v0.2.0

- [ ] Stream-like Operations: FlatMap, GroupBy, ...
  
### v0.3.0

- [ ] Bean: Basic bean, Bean copy, Bean compare
- [ ] Concurrency/Coroutine: CoroutinePool/WorkerPool ...

### v0.4.0

### v1.0.0 and later

- [ ] Advanced Functional Programming: Monad, Rx, ...
- [ ] Advanced Functional Programming: **Pattern Matching**
- [ ] Encapsulated Functions using Advanced Functional Programming Features

## Design Philosophy

### Functional Programming vs Go

#### Records of v0.0.1

- Nil slice or "Empty" slice when error occurs?
  - We all know that in Go, a default initialized slice is `nil`, and an empty slice is `[]T{}`, they are different.
  - Though some functional programming languages tend to use empty list to represent "nothing", we still prefer to use `nil` here.
  - As I believe that when an `error` occurs, the result should be abandoned, and the `nil` slice is a better choice.

- FP Features in Underlying Packages?
  - Maybe not. That would cost unnecessary time and space overhead, even if the Go compiler seems to do a pretty good job of optimizing for it.
  - In functional programming, we usually use `Option` or `Maybe` to represent a value that may be `nil` or `None`. But that was not implemented until v0.0.2.
  - Yet you can also use `functional.Some(slice.Insert(...))` to wrap the result.
  - We will also implement encapsulated functions like `slice.InsertOrError(...)` to return `Option` or `Maybe` in the future.(>=v0.4.0 or v1.0.0)

### Licensing

This project is licensed under the Mizumoto.General.Public.License - see the [LICENSE](./LICENSE) file.
As for the full context of this license, please refer to the markdown version: [Mizumoto General Public License v1.5](./licensing/Mizumoto.General.Public.License.v1.5.md).

---

copyRight @ 2024 Ruiyuan "mizumoto-cn" Xu <mizumoto@mizumoto.tech>
