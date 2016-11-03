# go-xmlrpc

##### Warning! This project is under heavy development!!
go-xml rpc takes another path to provide parsing of xml. Instead of reflect for binding xml => go, go-xmlrpc generates code for xml parsing.
It uses etree implementation and generates code directly for service methods.
Isn't that nice?
It's safe, fast, awesome!

Ok let's have a look at example

```go
//go:generate xmlrpcgen --file $GOFILE HelloService
package example

type HelloService struct {
    Config Config
}

func (h *HelloService) Search(query string, page int, isit bool) ([]string, error) {
    return []string{}, nil
}
```

That's pretty simple service with search method. Now when you run `go generate` go-xmlrpc generates file with
all parsing code.
go-xmlrpc parses your service methods and generates xml parsing code directly to your methods.
Nice way is that you can reuse this service in go code.

go-xmlrpc creates http handler for you

```go
handler := xmlrpc.Handler()
if err := handler.AddService(&HelloService{Config:Config}, "hello"); err != nil {
    panic(err)
}
```

You can then call methods `hello.Search` with your favorite xmlrpc client.
You can use then handler directly in your favorite mux router since it is Handler.

## Return values:

Your service methods must return either:
* error - simple error or xmlrpc.Error with code (`xmlrpc.Errorf(400, "this %s", "error)`)
* result and error

This is because xml rpc should return at least error.

## Error:
If you return xmlrpc error with added code it will be addedded to `result.fault`.
Otherwise error code will be 500.

## Features:

* since go-xmlrpc generates code in your package, you can use also unexported methods (yay)
* you can register your services with instantiated database connections, or other variables
* Automatically adds `system.listMethods` with all available methods
* inspect service method arguments and return values recursively (yay nice!)

## Limitations:

* currently arguments and return values can handle only inline definitions and not named arguments (not added in first version, will probably change in future)
* Registered services must be pointers (just to be sure all your methods are usable)

## Gotchas:

Don't forget to rerun `go generate` when you either:

* change definition of service methods
* remove service methods
* add service methods

## TODO:
* Add proper error messages to parse errors (with whole path). 
* Cleanup code generation with proper documentation
* Possibly remove temporary variables in parsing code.

## Contributions:
Your PRs are welcome

## Author:
phonkee
