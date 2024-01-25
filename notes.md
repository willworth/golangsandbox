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
