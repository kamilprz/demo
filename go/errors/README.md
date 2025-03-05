# Errors in Go

Credit to "Learning Go" by Jon Bodner for much of this information and some code snippets.

## What are Errors?

`error` is a built-in interface which provides a single method:

```go
type error interface {
    Error() string
}
```

Anything that implements this interface is considered an error.

Go handles errors by returning a value of type `error` as the last return for a function.

- This is entirely by convention, but it is one that is so strong that it should essentially never be breached.

The Go compiler requires all variables to be read. Making errors returned values forces developers to either check and handle error conditions or make it explicit that they are ignoring errors by using an underscore `_`.

When a function executes:

- As expected -> `nil` is returned as error parameter.
- Something goes wrong -> error value is returned.

It is the calling functions responsibility to check the error return value by comparing it to `nil`, handling the error or returning an error of its own. Since Go does not have try/catch blocks, this is handled by `if` statements.

### Printing Errors

There are two ways to pass strings and create an error:

- `errors.New()`
- `fmt.Errorf()`

Error messages should **not** be capitalized nor should they end with punctuation or a newline.

In most cases, you should set the other return values to their zero values when a non-nil error is returned.

The related code is under [`basics/main.go`](./basics/main.go).

## Sentinel Errors

These are errors which signal that processing cannot continue because of a problem with the current state.

Sentinel errors are one of the few variables that are declared at the package level. By convention, their names start with `Err`. They should be treated as read-only.

An example of a sentinel error can be found under [`sentinel/main.go`](./sentinel/main.go).

In general you should try to reuse existing Sentinel Errors as defining your own makes them part of your public API and must be maintained for backward compatibility.

See [`my_sentinel/main.go`](./my_sentinel/main.go) for an example of custom Sentinel Errors and ways of defining them.

## Extending Errors

Since `error` is an interface, you can define your own errors that include additional information for logging or error handling.

For example, you might want to include a status code as part of the error to indicate the kind of error that should be reported back to the user. This lets you avoid string comparisons (whose text might change) to determine error causes.

See [`my_status_err/main.go`](./my_status_err/main.go) for example code of this. By default this will not fail. See what happens when you enter an invalid `user` or `file`!

If you are using your own error type, be sure you don’t return an uninitialized instance. The code above also covers what happens if you don't.

## Error Wrapping

When you preserve an error while adding information, it is called wrapping the error When you have a series of wrapped errors, it is called an error tree. You can also unwrap (although you don't usually do this - use `Is` and `As` instead).

The `fmt.Errorf` function has a special verb, `%w` which can be used to create an error whose formatted string includes the formatted string of another error and which contains the original error as well. The convention is to write `: %w` at the end of the error format string and to make the error to be wrapped the last parameter passed to `fmt.Errorf`.

The best way to understand is to look at some code [`wrapping/main.go`](./wrapping/main.go).

If you want to wrap an error with your custom error type, your error type needs to implement the method `Unwrap`.

If a sentinel error is wrapped, you cannot use `==` to check for it.

If you want to create a new error that contains the message from another error, but don’t want to wrap it, use `fmt.Errorf` to create an error but use the `%v` verb instead of `%w`.

```go
err := internalFunction()
if err != nil {
    return fmt.Errorf("internal failure: %v", err)
}
```

You can implement your own error type that supports multiple wrapped errors by implementing `Unwrap` but have it return `[]error` instead of `error`. Note that Go doesn’t support method overloading, so you can’t create a single type that provides both implementations of `Unwrap`.

Due to the limitations of `Unwrap`, it is recommended to use `Is` and `As` instead.

## Is and As

tldr; Use `errors.Is` when you are looking for a specific instance or specific values. Use `errors.As` when you are looking for a specific type.

### errors.Is

To check whether the returned error or any errors that it wraps match a specific sentinel error instance, use `errors.Is`.

- returns `true` if any error in the error tree matches the provided sentinel error

By default, `errors.Is` uses `==` to compare each wrapped error with the specified error. If this does not work for an error type that you define (for example, if your error is a noncomparable type), implement the `Is` method on your error.

Another use for defining your own `Is` method is to allow comparisons against errors that aren’t identical instances. For example matching all errors related to Databases.

```go
if errors.Is(err, ResourceErr{Resource: "Database"}) {
    fmt.Println("The database is broken:", err)
    // process the codes
}
```

Check out [`is/main.go`](./is/main.go) for some code examples.

### errors.As

To check whether a returned error (or any error it wraps) matches a specific type, use `errors.As`.

- returns true if an error in the error tree was found that matched, and that matching error is assigned to the second parameter

Implementing an `As` method is non-trivial and should only be done in unusual circumstances, such as when you want to match an error of one type and return another.

```go
err := AFunctionThatReturnsAnError()
var myErr MyErr

var coder interface {
    CodeVals() []int
}

if errors.As(err, &myErr) {
    fmt.Println(myErr.Codes)
}

// the second parameter to errors.As can be a pointer to an interface
if errors.As(err, &coder) {
    fmt.Println(coder.CodeVals())
}
```

Check out [`as/main.go`](./as/main.go) for some code examples.

## Panic and Recover

A panic is a state generated by the Go runtime whenever it is unable to figure out what should happen next.

As soon as a panic happens, the current function exits immediately, and any defers attached to the current function start running. When those defers complete, the defers attached to the calling function run, and so on, until main is reached. The program then exits with a message and a stack trace.

If there is a panic in a goroutine other than the main goroutine, the chain of defers ends at the function used to launch the goroutine. A program exits if any goroutine panics without being recovered.

You can create your own panics using the built-in function `panic` which takes in one parameter of any type, typically a string.

See [`panic/main.go`](./panic/main.go) for an example.

There is a way to recover from panics using the built-in `recover` function which is called from within a defer to check whether a panic happened. If there was a panic, the value assigned to the panic is returned. Once a recover happens, execution continues normally.

There’s a specific pattern for using recover.

- You register a function with defer to handle a potential panic.
- You call recover within an if statement and check whether a non-nil value was found.
- You must call recover from within a defer because once a panic happens, only deferred functions are run.

See [`recover/main.go`](./recover/main.go) for an example.

Recover doesn't make it clear what could fail, only that if something does, execution can continue.

Reserve panics for fatal situations and use recover as a way to gracefully handle these situations. You will rarely want to keep the program running after a panic occurs - it's typically for a reason.

Panics should not escape the boundary of your public API. If necessary, recover and convert it to an error - allowing the calling code to decide how to handle the situation.
