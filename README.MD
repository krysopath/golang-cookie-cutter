# the golang cookie cutter

A naive and simple golang project to start out with.

# `go mod`

Go modules are the preferred way to bundle dependencies since golang 1.11.
After you have some code written, golang can create a module (two files with
package names, version and checksums)

We create the module in the beginning:

```
 $ go mod init github.com/krysopath/golang-cookie-cutter.git
go: creating new go.mod: module github.com/krysopath/golang-cookie-cutter.git
```

> This created a file `go.mod` naming all our dependencies and their versions.
> (We have no dependencies currently)

```
$ cat go.mod
module github.com/krysopath/golang-cookie-cutter.git

go 1.14
```

After we make changes to the dependecies, we want to run run `go mod tidy`,
which will prune/update `go.mod` with the new requirements.

> `go mod tidy` will also create the file `go.sum` if we have declared any
> dependencies outside of go's stdlib.


# go packages

A package is what is declared as requirements in your go module. Also your own
project is at least one package.

Go packages are strongly coupled with a git repository, and no official binary
distributor exists.

Some tried, but fairly quickly realized, that building golang on the fly is too
cheap and easy.

> go packages are stored in git for version control

> as long as golang can `git pull` on a remote, it can find go packages in there.

> A fully qualified golang package might look like this
> `git.domain.tld/namespace/project/package maj.min.pat`

# introduce a dependency

I want my cookies to be marshaled into yaml, so I require the official yaml library.

Downlod and compile this via:
```
go get gopkg.in/yaml.v2 v2.3.0
```

> This will save the compiled lib in `$GOPATH` and makes it available for
> linking (and source inspection) so we can use it in our cookie cutter.


Now, if we use a name/symbol from `yaml.v2`, we can automatically type hint and
complete our code. 



## assignments


```
package main

import "fmt"

func main() {

    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "apple"
    fmt.Println(f)
}``


### functions

A simple function
```
func SimpleFunction() {
    fmt.Println("Hello World")
}
```


A fucntion with arguments
```
func add(x int, y int) {
    total := 0
    total = x + y
    fmt.Println(total)
}
func sub(x, y int) {
    total := 0
    total = x + y
    fmt.Println(total)
}
```




A function with return value
```
// Function with int as return type
func add(x, y int) int {
    total := 0
    total = x + y
    return total
}

```

A function with named returns:
```
func rectangle(l int, b int) (area int) {
    var parameter int
    parameter = 2 * (l + b)
    fmt.Println("Parameter: ", parameter)

    area = l * b
    return 
}```


## inheritance?!

> Go doesn't have inheritance – instead composition, embedding and interfaces support code reuse and polymorphism.

:)


