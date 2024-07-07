# Functional Package Documentation 2 (Functional Programming Operations)

The `functional` package in `github.com/mizumoto-cn/fpkit` provides a set of helper methods for working with functional programming concepts in Go. This includes operations on functions, numeric types, and atomic operations.

## Import Method

```go
import "github.com/mizumoto-cn/fpkit/functional"
```

## Documentation

### functional/compose.go

1. **Compose**
   - **Description**: Returns a function that is the composition of the given functions from right to left.
   - **Input**: Variadic functions `fns` of type `func(T) T`.
   - **Output**: A composed function that applies the given functions in right-to-left order.
   - **Usage**:

     ```go
     f := functional.Compose(
         func(x int) int { return x + 1 },
         func(x int) int { return x * 2 },
     )
     result := f(3) // result: 7
     ```

2. **Pipe**
   - **Description**: Returns a function that is the composition of the given functions from left to right.
   - **Input**: Variadic functions `fns` of type `func(T) T`.
   - **Output**: A composed function that applies the given functions in left-to-right order.
   - **Usage**:

     ```go
     f := functional.Pipe(
         func(x int) int { return x + 1 },
         func(x int) int { return x * 2 },
     )
     result := f(3) // result: 8
     ```

### functional/curry.go

1. **Curry2**
   - **Description**: Returns a curried version of a function that takes 2 arguments.
   - **Input**: A function `fn` of type `func(A, B) R`.
   - **Output**: A curried function of type `func(A) func(B) R`.
   - **Usage**:

     ```go
     add := func(a, b int) int { return a + b }
     curriedAdd := functional.Curry2(add)
     result := curriedAdd(1)(2) // result: 3
     ```

2. **Curry3**
   - **Description**: Returns a curried version of a function that takes 3 arguments.
   - **Input**: A function `fn` of type `func(A, B, C) R`.
   - **Output**: A curried function of type `func(A) func(B) func(C) R`.
   - **Usage**:

     ```go
     add := func(a, b, c int) int { return a + b + c }
     curriedAdd := functional.Curry3(add)
     result := curriedAdd(1)(2)(3) // result: 6
     ```

3. **Curry4**
   - **Description**: Returns a curried version of a function that takes 4 arguments.
   - **Input**: A function `fn` of type `func(A, B, C, D) R`.
   - **Output**: A curried function of type `func(A) func(B) func(C) func(D) R`.
   - **Usage**:

     ```go
     add := func(a, b, c, d int) int { return a + b + c + d }
     curriedAdd := functional.Curry4(add)
     result := curriedAdd(1)(2)(3)(4) // result: 10
     ```

4. **Curry5**
   - **Description**: Returns a curried version of a function that takes 5 arguments.
   - **Input**: A function `fn` of type `func(A, B, C, D, E) R`.
   - **Output**: A curried function of type `func(A) func(B) func(C) func(D) func(E) R`.
   - **Usage**:

     ```go
     add := func(a, b, c, d, e int) int { return a + b + c + d + e }
     curriedAdd := functional.Curry5(add)
     result := curriedAdd(1)(2)(3)(4)(5) // result: 15
     ```

5. **Curry6**
   - **Description**: Returns a curried version of a function that takes 6 arguments.
   - **Input**: A function `fn` of type `func(A, B, C, D, E, F) R`.
   - **Output**: A curried function of type `func(A) func(B) func(C) func(D) func(E) func(F) R`.
   - **Usage**:

     ```go
     add := func(a, b, c, d, e, f int) int { return a + b + c + d + e + f }
     curriedAdd := functional.Curry6(add)
     result := curriedAdd(1)(2)(3)(4)(5)(6) // result: 21
     ```

### functional/filter.go

1. **Filter**
   - **Description**: Filters elements in a slice based on a predicate function.
   - **Input**: A predicate function `fn` of type `func(T, int) bool` and variadic input values of type `T`.
   - **Output**: A slice of values that satisfy the predicate.
   - **Usage**:

     ```go
     isEven := func(x int, _ int) bool { return x%2 == 0 }
     filtered := functional.Filter(isEven, 1, 2, 3, 4, 5)
     // filtered: [2, 4]
     ```

### functional/map.go

1. **Map**
   - **Description**: Applies a function to each element in a slice and returns a slice of the results.
   - **Input**: A function `f` of type `func(T) U` and variadic input values of type `T`.
   - **Output**: A slice of type `U` containing the results of applying the function to the input values.
   - **Usage**:

     ```go
     double := func(x int) int { return x * 2 }
     mapped := functional.Map(double, 1, 2, 3, 4, 5)
     // mapped: [2, 4, 6, 8, 10]
     ```

### functional/reduce.go

1. **Foldl**
   - **Description**: Applies a function to each element of a list, starting from the left, and returns the final value.
   - **Input**: A slice `s` of type `T`, a function `fn` of type `func(U, T) U`, and an initial value `init` of type `U`.
   - **Output**: The final value after applying the function to each element from left to right.
   - **Usage**:

     ```go
     sum := func(acc, x int) int { return acc + x }
     result := functional.Foldl([]int{1, 2, 3, 4}, sum, 0)
     // result: 10
     ```

2. **Foldr**
   - **Description**: Applies a function to each element of a list, starting from the right, and returns the final value.
   - **Input**: A slice `s` of type `T`, a function `fn` of type `func(U, T) U`, and an initial value `init` of type `U`.
   - **Output**: The final value after applying the function to each element from right to left.
   - **Usage**:

     ```go
     concat := func(acc string, x int) string { return acc + strconv.Itoa(x) }
     result := functional.Foldr([]int{1, 2, 3, 4}, concat, "")
     // result: "4321"
     ```

3. **Reduce**
   - **Description**: Applies a reduction function to each element of the slice on a left-to-right basis to a given initial value.
   - **Input**: A slice `s` of type `T`, a function `fn` of type `func(U, T) U`, and an initial value `init` of type `U`.
   - **Output**: The final value after applying the function to each element from left to right.
   - **Usage**:

     ```go
     product := func(acc, x int) int { return acc * x }
     result := functional.Reduce([]int{1, 2, 3, 4}, product, 1)
     // result: 24
     ```