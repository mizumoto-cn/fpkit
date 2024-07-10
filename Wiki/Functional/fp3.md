# Functional Package Documentation 3 (Optional/Maybe Types)

The `functional` package in `github.com/mizumoto-cn/fpkit` provides a set of helper methods for working with functional programming concepts in Go. This includes operations on functions, numeric types, and atomic operations.

## Import Method

```go
import "github.com/mizumoto-cn/fpkit/functional"
```

## Functions

### functional/optional.go

1. **Optional Interface**
   - **Description**: Defines an interface for optional values with various utility methods.
   - **Methods**:
     - `IsPresent() bool`: Returns `true` if the value is present, otherwise `false`.
     - `IsNil() bool`: Returns `true` if the value is not present, otherwise `false`.
     - `IsValid() bool`: Returns `true` if the value is valid.
     - `IsPtr() bool`: Returns `true` if the value is a pointer.
     - `Just(any) Optional[any]`: Returns an optional with the given value.
     - `OrElse(T) T`: Returns the value if present, otherwise returns the provided default value.
     - `Clone() Optional[T]`: Returns a clone of the optional value.
     - `FlatMap(func(T) Optional[T]) Optional[T]`: Applies a function to the value if present and returns the result.
     - `IfPresent(func())`: Executes a function if the value is present.
     - `Kind() reflect.Kind`: Returns the kind of the value.
     - `Type() reflect.Type`: Returns the type of the value.
     - `IsKindOf(reflect.Kind) bool`: Checks if the value is of the specified kind.
     - `IsTypeOf(reflect.Type) bool`: Checks if the value is of the specified type.
     - `Unwrap() T`: Returns the value.
     - `UnwrapAny() any`: Returns the value as an `any`.

2. **maybe Struct**
   - **Description**: Implements the `Optional` interface for generic types.
   - **Fields**:
     - `value T`: The value.
     - `isNil bool`: Indicates if the value is nil.

3. **none Struct**
   - **Description**: Implements the `Optional` interface for representing absence of a value.
   - **Fields**:
     - `maybe[any]`: Embeds a `maybe` struct with `any` type.

4. **Just**
   - **Description**: Creates an `Optional` with the given value.
   - **Input**: A value of any type `T`.
   - **Output**: An `Optional` containing the value.
   - **Usage**:

     ```go
     opt := functional.Just(42)
     ```

5. **MakeClone**
   - **Description**: Creates a clone of the `Optional` object.
   - **Input**: An `Optional` object and a destination pointer.
   - **Output**: A cloned `Optional` object.
   - **Usage**:

     ```go
     clone := functional.MakeClone(opt, new(int))
     ```

6. **IsPresent**
   - **Description**: Checks if the value is present.
   - **Output**: Returns `true` if the value is present, otherwise `false`.
   - **Usage**:

     ```go
     present := opt.IsPresent()
     ```

7. **IsNil**
   - **Description**: Checks if the value is nil.
   - **Output**: Returns `true` if the value is nil, otherwise `false`.
   - **Usage**:

     ```go
     isNil := opt.IsNil()
     ```

8. **IsValid**
   - **Description**: Checks if the value is valid.
   - **Output**: Returns `true` if the value is valid, otherwise `false`.
   - **Usage**:

     ```go
     isValid := opt.IsValid()
     ```

9. **IsPtr**
   - **Description**: Checks if the value is a pointer.
   - **Output**: Returns `true` if the value is a pointer, otherwise `false`.
   - **Usage**:

     ```go
     isPtr := opt.IsPtr()
     ```

10. **OrElse**
    - **Description**: Returns the value if present, otherwise returns the default value.
    - **Input**: A default value of type `T`.
    - **Output**: The value if present, otherwise the default value.
    - **Usage**:

      ```go
      value := opt.OrElse(0)
      ```

11. **Clone**
    - **Description**: Returns a clone of the optional value.
    - **Output**: A cloned `Optional` object.
    - **Usage**:

      ```go
      clone := opt.Clone()
      ```

12. **FlatMap**
    - **Description**: Applies a function to the value if present and returns the result.
    - **Input**: A function `fn` of type `func(T) Optional[T]`.
    - **Output**: The result of applying the function to the value if present.
    - **Usage**:

      ```go
      result := opt.FlatMap(func(x int) functional.Optional[int] {
          return functional.Just(x * 2)
      })
      ```

13. **IfPresent**
    - **Description**: Executes a function if the value is present.
    - **Input**: A function `fn` of type `func()`.
    - **Usage**:

      ```go
      opt.IfPresent(func() {
          fmt.Println("Value is present")
      })
      ```

14. **Kind**
    - **Description**: Returns the kind of the value.
    - **Output**: The `reflect.Kind` of the value.
    - **Usage**:

      ```go
      kind := opt.Kind()
      ```

15. **Type**
    - **Description**: Returns the type of the value.
    - **Output**: The `reflect.Type` of the value.
    - **Usage**:

      ```go
      typ := opt.Type()
      ```

16. **IsKindOf**
    - **Description**: Checks if the value is of the specified kind.
    - **Input**: A `reflect.Kind`.
    - **Output**: Returns `true` if the value is of the specified kind, otherwise `false`.
    - **Usage**:

      ```go
      isKind := opt.IsKindOf(reflect.Int)
      ```

17. **IsTypeOf**
    - **Description**: Checks if the value is of the specified type.
    - **Input**: A `reflect.Type`.
    - **Output**: Returns `true` if the value is of the specified type, otherwise `false`.
    - **Usage**:

      ```go
      isType := opt.IsTypeOf(reflect.TypeOf(42))
      ```

18. **Unwrap**
    - **Description**: Returns the value.
    - **Output**: The value.
    - **Usage**:

      ```go
      value := opt.Unwrap()
      ```

19. **UnwrapAny**
    - **Description**: Returns the value as `any`.
    - **Output**: The value as `any`.
    - **Usage**:

      ```go
      value := opt.UnwrapAny()
      ```

### Package Variables

1. **Maybe**
   - **Description**: A package variable representing an empty `Optional`.
   - **Type**: `Optional[any]`
   - **Usage**:

     ```go
     var maybeValue functional.Optional[any] = functional.Maybe.Just(42)
     ```

2. **None**
   - **Description**: A package variable representing an absent value.
   - **Type**: `Optional[any]`
   - **Usage**:

     ```go
     var noneValue functional.Optional[any] = functional.None
     ```
