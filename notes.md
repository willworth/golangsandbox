https://go.dev/doc/effective_go
https://gobyexample.com/
https://go.dev/doc/tutorial/
https://pkg.go.dev/std

packages are made up of modules

https://github.com/mschwarzmueller/golang-course-code

every go file must be part of a package.

if you have both files as part of eg main package, no import statement is required. functions are shared across files, but imports aren´t. eg fmt.

every package must have its own subfolder

only upper case funcs and vars are available in other packages.

https://pkg.go.dev/github.com/pallinder/go-randomdata

go get github.com/Pallinder/go-randomdata

go get is the global install

# pointers

variables that store value addresses instead of values

why pointers

1. avoid unnecessary value copies
2. directly mutate values

dereferencing = \*varName = gives the value of the address stored as a pointer

you can't perform pointer maths

pointers to structs don't need to be dereferenced.

with anonymous embedding you can access the fields and methods directly on the struct

scan is mostly used for single word input

An interface is a contract that guarantees a val(normally a struct) has a certain method. They don't define logic, just state that the method and its logic exists somewhere
save => saver if only contains one method is the convention

## Generics

Generics in Go are used to write flexible, reusable functions and types that can operate on different types while still providing type safety. They are particularly useful when you want to create a function or a data structure that can work with any data type, not just a specific one.

The T in square brackets ([T]) is a type parameter. It's a placeholder for any type that will be specified when the function is called. The name T is conventional, but you can use any valid identifier.

```go
func add[T any](a,b T)(T, T){
    return a + b
}


```

## maps

anything as a key
structs are predefined structures
if it's a custom collection with labels, you want a map
