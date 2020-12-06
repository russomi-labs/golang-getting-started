# golang-getting-started

Getting Started tutorials for Go

## Tutorial: Get started with Go

In this tutorial, you'll get a brief introduction to Go programming. Along the way, you will:

- Install Go (if you haven't already).
- Write some simple "Hello, world" code.
- Use the go command to run your code.
- Use the Go package discovery tool to find packages you can use in your own code.
- Call functions of an external module.

### Install Go

- [Download and install](https://golang.org/doc/install)
- Install via Brew on macOS:

``` bash

# output a list of what you can install
brew search golang

# Go ahead and install it
brew install golang

# check the version of Go
go version

# to update Go, you can run the following commands
brew update
brew upgrade golang

```

### Write some code

Create a hello directory for your first Go source code.

``` BASH
mkdir hello
cd hello
```

In your text editor, create a file hello.go in which to write your code.

``` Go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Run your code to see the greeting.

``` BASH
$ go run hello.go

Hello, World!
```

``` BASH
$ go help

```

### Call code in an external package

When you need your code to do something that might have been implemented by someone else, you can look for a package that has functions you can use in your code.

- Visit [pkg.go.dev](https://pkg.go.dev) and search for a "quote" package.
- Locate and click the `rsc.io/quote` package in search results (if you see rsc.io/quote/v3, ignore it for now).

``` Go
package main

import "fmt"

import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}
```

- Put your own code in a module for tracking dependencies.

``` BASH
go mod init hello
```

- Run your code to see the message generated by the function you are calling.

``` BASH
$ go run hello.go

go: finding module for package rsc.io/quote
go: found rsc.io/quote in rsc.io/quote v1.5.2
Don't communicate by sharing memory, share memory by communicating.
```

## Tutorial: Create a Go module

In this [tutorial](https://golang.org/doc/tutorial/create-module) you'll create two modules.

- The first is a library which is intended to be imported by other libraries or applications.
- The second is a caller application which will use the first.

This tutorial's sequence includes six brief topics that each illustrate a different part of the language.

- Create a module -- Write a small module with functions you can call from another module.
- Call your code from another module -- Import and use your new module.
- Return and handle an error -- Add simple error handling.
- Return a random greeting -- Handle data in slices (Go's dynamically-sized arrays).
- Return greetings for multiple people -- Store key/value pairs in a map.
- Add a test -- Use Go's built-in unit testing features to test your code.
- Compile and install the application -- Compile and install your code locally.

### Create a module

``` BASH
# Create a greetings directory for your Go module source code.
# This is where you'll write your module code.
mkdir greetings
cd greetings

# Start your module using the go mod init command to create a go.mod file.
go mod init github.com/russomi-labs/golang-getting-started/greetings
```

Create a file in which to write your code and call it `greetings.go` .

``` Go
package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
```

This is the first code for your module.  It returns a greeting to any caller that asks for one.

In this code, you:

- Declare a greetings package to collect related functions.
- Implement a Hello function to return the greeting.
- This function takes a name parameter whose type is string, and returns a string.
- In Go, a function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an [exported name](https://tour.golang.org/basics/3).
- Declare a message variable to hold your greeting.
- In Go, the := operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type).
- Taking the long way, you might have written this as:

``` Go
var message string
message = fmt. Sprintf("Hi, %v. Welcome!", name)
```

- Use the `fmt` package's `Sprintf` function to create a greeting message.
- The first argument is a format string, and `Sprintf` substitutes the name parameter's value for the `%v` format verb.
- Inserting the value of the name parameter completes the greeting text.
- Return the formatted greeting text to the caller.

### Call your code from another module

Create a hello directory for your Go module source code. This is where you'll write your caller.

``` BASH
cd ..
mkdir hello
cd hello

touch hello.go
```

``` Go
package main

import (
    "fmt"

    "github.com/russomi-labs/golang-getting-started/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
```

Create a new module for this hello package.

``` BASH
go mod init hello
```

Edit the hello module to use the unpublished greetings module.

``` Go
module hello

go 1.14

replace github.com/russomi-labs/golang-getting-started/greetings => ../greetings
```

Here, the [replace directive](https://golang.org/ref/mod#tmp_15) tells Go to replace the module path (the URL example.com/greetings) with a path you specify. In this case, that's a greetings directory next to the hello directory.

In the hello directory, run go build to make Go locate the module and add it as a dependency to the go.mod file.

``` BASH
go build
```

Look at `go.mod` again to see the changes made by `go build` , including the `require` directive Go added.

``` Go
module hello

go 1.15

replace github.com/russomi-labs/golang-getting-started/greetings => ../greetings

require github.com/russomi-labs/golang-getting-started/greetings v0.0.0-00010101000000-000000000000

```

In the hello directory, run the hello executable (created by go build) to confirm that the code works.

``` BASH
$ ./hello
Hi, Gladys. Welcome!
```

### Return and handle an error

Handling errors is an essential feature of solid code. In this section, you'll add a bit of code to return an error from the greetings module, then handle it in the caller.

There's no sense sending a greeting back if you don't know who to greet. Return an error to the caller if the name is empty. Copy the following code into greetings.go and save the file.

``` Go
package greetings

import (
    "errors"
    "fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name")
    }

    // If a name was received, return a value that embeds the name
    // in a greeting message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}
```

- Change the function so that it returns two values: a string and an error.
- Your caller will check the second value to see if an error occurred. (Any Go function can return [multiple values](https://golang.org/doc/effective_go.html#multiple-returns).)
- Import the Go standard library errors package so you can use its [errors. New function](https://golang.org/pkg/errors/#example_New).
- Add an if statement to check for an invalid request and return an error if the request is invalid.
- The errors. New function returns an error with your message inside.
- Add nil (meaning no error) as a second value in the successful return.
- That way, the caller can see that the function succeeded.

In your hello/hello.go file, handle the error now returned by the Hello function, along with the non-error value.

``` Go
package main

import (
    "fmt"
    "log"

    "github.com/russomi-labs/golang-getting-started/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Request a greeting message.
    message, err := greetings.Hello("")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}
```

- Configure the [log package](https://golang.org/pkg/log/) to print the command name ("greetings: ") at the start of its log messages, without a time stamp or source file information.
- Assign both of the Hello return values, including the error, to variables.
- Change the Hello argument from Gladys's name to an empty string, so you can try out your error-handling code.
- Look for a non-nil error value. There's no sense continuing in this case.
- Use the functions in the standard library's log package to output error information.
- If you get an error, you use the log package's [Fatal function](https://pkg.go.dev/log?tab=doc#Fatal) to print the error and stop the program.

At the command line in the hello directory, run hello.go to confirm that the code works.

``` BASH
$ go run hello.go
greetings: empty name
exit status 1
```

### Return a random greeting

In this section, you'll change your code so that instead of returning a single greeting every time, it returns one of several predefined greeting messages.

- To do this, you'll use a Go [slice](https://blog.golang.org/slices-intro).
- A slice is like an array, except that it's dynamically sized as you add and remove items. It's one of the most useful types in Go.
- You'll add a small slice to contain three greeting messages, then have your code return one of the messages randomly.

In greetings/greetings.go, change your code so it looks like the following.

``` Go
// Gopackage greetings

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// init sets initial values for variables used in the function.
func init() {
    rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}

```

In this code, you:

- Add a randomFormat function that returns a randomly selected format for a greeting message.
- Note that randomFormat starts with a lowercase letter, making it accessible only to code in its own package (in other words, it's not exported).
- In randomFormat, declare a formats slice with three message formats. When declaring a slice, you omit its size in the brackets, like this: []string. This tells Go that the array underlying a slice can be dynamically sized.
- Use the [math/rand package](https://golang.org/pkg/math/rand/) to generate a random number for selecting an item from the slice.
- Add an [init function](https://golang.org/doc/effective_go.html#init) to seed the rand package with the current time. Go executes init functions automatically at program startup, after global variables have been initialized.
- In Hello, call the randomFormat function to get a format for the message you'll return, then use the format and name value together to create the message.
- Return the message (or an error) as you did before.

At the command line, change to the hello directory, then run hello.go to confirm that the code works. Run it multiple times, noticing that the greeting changes.

Oh -- don't forget to add Gladys's name (or a different name, if you like) as an argument to the Hello function call in hello.go: greetings. Hello("Gladys")

``` BASH
$ go build
$ ./hello
Great to see you, Gladys!

$ ./hello
Hi, Gladys. Welcome!

$ ./hello
Hail, Gladys! Well met!

```

### Return greetings for multiple people

In the last changes you'll make to your module's code, you'll add support for getting greetings for multiple people in one request. In other words, you'll handle a multiple-value input and pair values with a multiple-value output.

To do this, you'll need to pass a set of names to a function that can return a greeting for each of them. Changing the Hello function's parameter from a single name to a set of names would change the function signature. If you had already published the greetings module and users had already written code calling Hello, that change would break their programs. In this situation, a better choice is to give new functionality a new name.

In the last code you'll add with this tutorial, update the code as if you've already published a version of the greetings module. Instead of changing the Hello function, add a new function Hellos that takes a set of names. Then, for the sake of simplicity, have the new function call the existing one. Keeping both functions in the package leaves the original for existing callers (or future callers who only need one greeting) and adds a new one for callers that want the expanded functionality.

In greetings/greetings.go, change your code so it looks like the following:

``` Go
package greetings

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
    // A map to associate names with messages.
    messages := make(map[string]string)
    // Loop through the received slice of names, calling
    // the Hello function to get a message for each name.
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        // In the map, associate the retrieved message with
        // the name.
        messages[name] = message
    }
    return messages, nil
}

// Init sets initial values for variables used in the function.
func init() {
    rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return one of the message formats selected at random.
    return formats[rand.Intn(len(formats))]
}
```

In this code, you:

- Add a Hellos function whose parameter is a slice of names rather than a single name. Also, you change one of its return types from a string to a map so you can return names mapped to greeting messages.
- Have the new Hellos function call the existing Hello function. This leaves both functions in place.
- Create a messages [map](https://blog.golang.org/maps) to associate each of the received names (as a key) with a generated message (as a value). In Go, you initialize a map with the following syntax: `make(map[key-type]value-type)` . You have the Hello function return this map to the caller.
- Loop through the names your function received, checking that each has a non-empty value, then associate a message with each. In this for loop, range returns two values: the index of the current item in the loop and a copy of the item's value. You don't need the index, so you use the Go [blank identifier (an underscore)](https://golang.org/doc/effective_go.html#blank) to ignore it.

In your `hello/hello.go` calling code, pass a slice of names, then print the contents of the names/messages map you get back.

In `hello.go` , change your code so it looks like the following.

``` Go
package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // A slice of names.
    names := []string{"Gladys", "Samantha", "Darrin"}

    // Request greeting messages for the names.
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    // If no error was returned, print the returned map of
    // messages to the console.
    fmt.Println(messages)
}

```

With these changes, you:

- Create a names variable as a slice type holding three names.
- Pass the names variable as the argument to the Hellos function.

At the command line, change to the directory that contains hello/hello.go, then run hello.go to confirm that the code works.
The output should be a string representation of the map associating names with messages, something like the following:

``` BASH
$ go run hello.go

map[Darrin:Hail, Darrin! Well met! Gladys:Hi, Gladys. Welcome! Samantha:Hail, Samantha! Well met!]
```

This topic introduced maps for representing name/value pairs. It also introduced the idea of [preserving backward compatibility](https://blog.golang.org/module-compatibility) by implementing a new function for new or changed functionality in a module. In the tutorial's next topic, you'll use built-in features to create a unit test for your code.

### Add a test

Go's built-in support for unit testing makes it easier to test as you go. Specifically, using naming conventions, Go's testing package, and the go test command, you can quickly write and execute tests.

In the greetings directory, create a file called `greetings_test.go` .

📓 Ending a file's name with `_test.go` tells the go test command that this file contains test functions.

In `greetings_test.go` , paste the following code and save the file.

``` Go
package greetings

import (
    "testing"
    "regexp"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile( `\b` +name+ `\b` )
    msg, err := Hello("Gladys")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf( `Hello("Gladys") = %q, %v, want match for %#q, nil` , msg, err, want)
    }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf( `Hello("") = %q, %v, want "", error` , msg, err)
    }
}
```

In this code, you:

- Implement test functions in the same package as the code you're testing.

- Create two test functions to test the greetings. Hello function. Test function names have the form TestName, where Name is specific to the test. Also, test functions take a pointer to the [testing package](https://golang.org/pkg/testing/)'s testing. T as a parameter. - You use this parameter's methods for reporting and logging from your test.

- Implement two tests:
  + TestHelloName calls the Hello function, passing a name value with which the function should be able to return a valid response message. If the call returns an error or an unexpected response message (one that doesn't include the name you passed in), you use the t parameter's [Fatalf](https://golang.org/pkg/testing/#T. Fatalf) method to print a message to the console and end execution.
  + TestHelloEmpty calls the Hello function with an empty string. This test is designed to confirm that your error handling works. If the call returns a non-empty string or no error, you use the t parameter's Fatalf method to print a message to the console and end execution.

At the command line in the greetings directory, run the [go test command](https://golang.org/cmd/go/#hdr-Test_packages) to execute the test.

The go test command executes test functions (whose names begin with Test) in test files (whose names end with _test.go). You can add the -v flag to get verbose output that lists all of the tests and their results.

The tests should pass.

``` BASH
$ go test
PASS
ok      example.com/greetings   0.364s

$ go test -v
=== RUN   TestHelloName
--- PASS: TestHelloName (0.00s)
=== RUN   TestHelloEmpty
--- PASS: TestHelloEmpty (0.00s)
PASS
ok      example.com/greetings   0.372s
```

Break the `greetings.Hello` function to view a failing test.

The `TestHelloName` test function checks the return value for the name you specified as a `Hello` function parameter. To view a failing test result, change the `greetings.Hello` function so that it no longer includes the name.

In `greetings/greetings.go` , paste the following code in place of the Hello function. Note that the highlighted lines change the value that the function returns, as if the name argument had been accidentally removed.

``` Go
// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    // message := fmt.Sprintf(randomFormat(), name)
    message := fmt.Sprint(randomFormat())
    return message, nil
}
```

At the command line in the greetings directory, run go test to execute the test.

This time, run go test without the -v flag. The output will include results for only the tests that failed, which can be useful when you have a lot of tests. The TestHelloName test should fail -- TestHelloEmpty still passes.

``` BASH
$ go test

--- FAIL: TestHelloName (0.00s)
    greetings_test.go:15: Hello("Gladys") = "Hail, %v! Well met!", <nil>, want match for `\bGladys\b` , nil
FAIL
exit status 1

FAIL    example.com/greetings   0.182s
```

This topic introduced Go's built-in support for unit testing.

### Compile and install the application

## References

- [How To Install Go and Set Up a Local Programming Environment on macOS](https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-macos)
- [Call code in an external package](https://golang.org/doc/tutorial/getting-started#call)

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you **would** like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
