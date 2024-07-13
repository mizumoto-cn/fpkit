# Functional Package Documentation 1 (Numerics, Reflection, and Atomic Operations)

The `functional` package in `github.com/mizumoto-cn/fpkit` provides a set of helper methods for working with functional programming concepts in Go. This includes operations on functions, numeric types, and atomic operations.

## Import Method

```go
import "github.com/mizumoto-cn/fpkit/functional"
```

## Documentation

### functional/types.go

1. **FnObject**
   - **Description**: Represents a function that takes any type as input and returns any type as output.
   - **Type**: `type FnObject func(...any) any`

2. **FnObj**
   - **Description**: Represents a function that takes input of type `T` and returns a single output of type `R`.
   - **Type**: `type FnObj[T, R any] func(...T) R`

### functional/reflect.go

1. **PtrOf**
   - **Input**: A value of any type `T`.
   - **Output**: A pointer to a copy of the passed-in value.
   - **Usage**:

     ```go
     ptr := functional.PtrOf(10)
     // ptr: *int pointing to 10
     ```

2. **SliceOf**
   - **Input**: Variadic values of any type `T`.
   - **Output**: A slice containing the passed-in values.
   - **Usage**:

     ```go
     slice := functional.SliceOf(1, 2, 3)
     // slice: []int{1, 2, 3}
     ```

3. **IsPtr**
   - **Input**: A value of any type `T`.
   - **Output**: Returns `true` if the value is a pointer, otherwise `false`.
   - **Usage**:

     ```go
     isPtr := functional.IsPtr(&struct{}{})
     // isPtr: true
     ```

4. **Kind**
   - **Input**: A value of any type `T`.
   - **Output**: Returns the `reflect.Kind` of the value.
   - **Usage**:

     ```go
     kind := functional.Kind(10)
     // kind: reflect.Int
     ```

5. **IsNil**
   - **Input**: A value of any type `T`.
   - **Output**: Returns `true` if the value is nil, otherwise `false`.
   - **Usage**:

     ```go
     isNil := functional.IsNil((*int)(nil))
     // isNil: true
     ```

### functional/numerics.go

1. **Sum**
   - **Input**: Variadic numeric values of type `T`.
   - **Output**: The sum of the numeric values.
   - **Usage**:

     ```go
     total := functional.Sum(1, 2, 3, 4)
     // total: 10
     ```

2. **CompareTo**
   - **Input**: Two orderable values `a` and `b` of type `T`.
   - **Output**: Returns `1` if `a` > `b`, `-1` if `a` < `b`, and `0` if `a` == `b`.
   - **Usage**:

     ```go
     comparison := functional.CompareTo(5, 3)
     // comparison: 1
     ```

3. **Sort**
   - **Input**: A slice of orderable values `a` of type `T` and a comparator function `cmp`.
   - **Output**: Sorts the slice in place based on the comparator.
   - **Usage**:

     ```go
     nums := []int{5, 3, 4, 1, 2}
     functional.Sort(nums, func(a, b int) bool { return a < b })
     // nums: [1, 2, 3, 4, 5]
     ```

4. **SortOrdered**
   - **Input**: A boolean `ascending` and variadic orderable values `input`.
   - **Output**: Returns a sorted slice of the input values in the specified order.
   - **Usage**:

     ```go
     sorted := functional.SortOrdered(true, 3, 1, 4, 1, 5, 9)
     // sorted: [1, 1, 3, 4, 5, 9]
     ```

5. **SortAsc**
   - **Input**: Variadic orderable values `input`.
   - **Output**: Returns a slice of the input values sorted in ascending order.
   - **Usage**:

     ```go
     sorted := functional.SortAsc(3, 1, 4, 1, 5, 9)
     // sorted: [1, 1, 3, 4, 5, 9]
     ```

6. **SortDesc**
   - **Input**: Variadic orderable values `input`.
   - **Output**: Returns a slice of the input values sorted in descending order.
   - **Usage**:

     ```go
     sorted := functional.SortDesc(3, 1, 4, 1, 5, 9)
     // sorted: [9, 5, 4, 3, 1, 1]
     ```

### functional/numerics_types.go

1. **Real**
   - **Description**: Defines a set of real numeric types including integers and floating-point numbers.
   - **Type**: `type Real interface { ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 }`

2. **Numeric**
   - **Description**: Extends `Real` to include complex numbers.
   - **Type**: `type Numeric interface { Real | ~complex64 | ~complex128 }`

3. **Orderable**
   - **Description**: Defines a set of types that can be ordered, including numeric types and strings.
   - **Type**: `type Orderable interface { ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string }`

4. **Comparator**
   - **Description**: Represents a comparator function for orderable types.
   - **Type**: `type Comparator[T Orderable] func(T, T) bool`

### functional/atomic.go

1. **AtomicBool**
   - **Description**: A struct that provides atomic operations for a boolean value using `sync/atomic`.
   - **Fields**: `value int32`
   - **Methods**:
     - **Set**
       - **Input**: A boolean value.
       - **Output**: Sets the atomic boolean to the given value.
       - **Usage**:

         ```go
         var ab functional.AtomicBool
         ab.Set(true)
         ```

     - **Get**
       - **Output**: Returns the current value of the atomic boolean.
       - **Usage**:

         ```go
         value := ab.Get()
         // value: true
         ```