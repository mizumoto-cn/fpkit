# Slice Package Documentation

Slices being the most commonly used data structure in Go, we provide a very large number of helper methods for it. All of the following methods are defined in the "github.com/mizumoto-cn/fpkit/slice" package.

## Import Method

```go
import "github.com/mizumoto-cn/fpkit/slice"
```

## Functions

1. **Contains**
   - **Input**: A slice `s` of type `T` and a value `v` of type `T`.
   - **Output**: Returns `true` if the slice contains the value, otherwise `false`.
   - **Usage**:

     ```go
     s := []int{1, 2, 3, 4, 5}
     slice.Contains(s, 3) // true
     ```

2. **ContainsFunc**
   - **Input**: A slice `s` of type `T` and a predicate function `eq` of type `func(T) bool`.
   - **Output**: Returns `true` if the slice contains an element that satisfies the predicate, otherwise `false`.
   - **Usage**:

     ```go
     s := []int{1, 2, 3, 4, 5}
     slice.ContainsFunc(s, func(x int) bool { return x > 3 }) // true
     ```

3. **ContainsAny**
   - **Input**: A slice `s` of type `T` and a variadic list of values `vs` of type `T`.
   - **Output**: Returns `true` if the slice contains any of the values, otherwise `false`.
   - **Usage**:

     ```go
     s := []int{1, 2, 3, 4, 5}
     slice.ContainsAny(s, 3, 7) // true
     ```

4. **ContainsAll**
   - **Input**: Two slices `mainSlice` and `subset` of type `T`.
   - **Output**: Returns `true` if all elements of the subset are in the main slice, otherwise `false`.
   - **Usage**:

     ```go
     mainSlice := []int{1, 2, 3, 4, 5}
     subset := []int{2, 3}
     slice.ContainsAll(mainSlice, subset) // true
     ```

5. **Delete**
   - **Input**: A slice `src` of type `T` and an index `index` of type `int`.
   - **Output**: Returns a new slice with the element at the given index removed, or an error if the index is out of range.
   - **Usage**:

     ```go
     src := []int{1, 2, 3, 4, 5}
     newSlice, err := slice.Delete(src, 2)
     // newSlice: [1, 2, 4, 5], err: nil
     ```

6. **DeleteMatched**
   - **Input**: A slice `src` of type `T` and a predicate function `match` of type `func(T) bool`.
   - **Output**: Returns a new slice with all elements that match the predicate removed.
   - **Usage**:

     ```go
     src := []int{1, 2, 3, 4, 5}
     newSlice := slice.DeleteMatched(src, func(x int) bool { return x%2 == 0 })
     // newSlice: [1, 3, 5]
     ```

7. **Difference**
   - **Input**: Two slices `s1` and `s2` of type `T`.
   - **Output**: Returns a slice of elements that are in `s1` but not in `s2`.
   - **Usage**:

     ```go
     s1 := []int{1, 2, 3, 4, 5}
     s2 := []int{3, 4, 5}
     diff := slice.Difference(s1, s2)
     // diff: [1, 2]
     ```

8. **Index**
   - **Input**: A slice `src` of type `T` and a value `v` of type `T`.
   - **Output**: Returns the index of the first occurrence of the value in the slice, or -1 if not present.
   - **Usage**:

     ```go
     src := []int{1, 2, 3, 4, 5}
     index := slice.Index(src, 3)
     // index: 2
     ```

9. **IndexMatchFunc**
   - **Input**: A slice `src` of type `T` and a predicate function `match` of type `func(T) bool`.
   - **Output**: Returns the index of the first element in the slice that satisfies the predicate, or -1 if not present.
   - **Usage**:

     ```go
     src := []int{1, 2, 3, 4, 5}
     index := slice.IndexMatchFunc(src, func(x int) bool { return x > 3 })
     // index: 3
     ```

10. **IndexAllMatchFunc**
    - **Input**: A slice `src` of type `T` and a predicate function `match` of type `func(T) bool`.
    - **Output**: Returns a slice of indexes of all elements in the slice that satisfy the predicate.
    - **Usage**:

      ```go
      src := []int{1, 2, 3, 4, 5}
      indexes := slice.IndexAllMatchFunc(src, func(x int) bool { return x%2 == 0 })
      // indexes: [1, 3]
      ```

11. **IndexAll**
    - **Input**: A slice `src` of type `T` and a value `v` of type `T`.
    - **Output**: Returns a slice of indexes of all occurrences of the value in the slice.
    - **Usage**:

      ```go
      src := []int{1, 2, 2, 3, 4, 2}
      indexes := slice.IndexAll(src, 2)
      // indexes: [1, 2, 5]
      ```

12. **LastIndexMatchFunc**
    - **Input**: A slice `src` of type `T` and a predicate function `match` of type `func(T) bool`.
    - **Output**: Returns the index of the last element in the slice that satisfies the predicate, or -1 if not present.
    - **Usage**:

      ```go
      src := []int{1, 2, 3, 4, 5}
      index := slice.LastIndexMatchFunc(src, func(x int) bool { return x%2 == 0 })
      // index: 3
      ```

13. **LastIndex**
    - **Input**: A slice `src` of type `T` and a value `v` of type `T`.
    - **Output**: Returns the index of the last occurrence of the value in the slice, or -1 if not present.
    - **Usage**:

      ```go
      src := []int{1, 2, 3, 4, 2, 5}
      index := slice.LastIndex(src, 2)
      // index: 4
      ```

14. **Insert**
    - **Input**: A slice `s` of type `T`, an index `index` of type `int`, and a value `value` of type `T`.
    - **Output**: Returns a new slice with the value inserted at the given index, or an error if the index is out of range.
    - **Usage**:

      ```go
      src := []int{1, 2, 4, 5}
      newSlice, err := slice.Insert(src, 2, 3)
      // newSlice: [1, 2, 3, 4, 5], err: nil
      ```
