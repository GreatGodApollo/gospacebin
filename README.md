<h1 align="center">gospacebin</h1>
<p align="center"><i>Made with :heart: by <a href="https://github.com/GreatGodApollo">@GreatGodApollo</a></i></p>
[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)][docs]

[Spacebin](https://spaceb.in) GoLang API library

## Built With
- [net/http](https://golang.org/pkg/net/http/)

## Usage
Import the package into your project
```go
import "github.com/GreatGodApollo/gospacebin"
```
Then construct a new spacebin client that can be used to access the API
```go
spacebin := gospacebin.NewClient("https://api.spaceb.in")
```
For more information on using the client read the [documentation](docs).
## Licensing

This project is licensed under the [MIT License](https://choosealicense.com/licenses/mit/)

## Authors

* [Brett Bender](https://github.com/GreatGodApollo)

[docs]: https://pkg.go.dev/github.com/GreatGodApollo/gospacebin