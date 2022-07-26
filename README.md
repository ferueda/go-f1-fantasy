go-f1-fantasy
=======


[![go-github release (latest SemVer)](https://img.shields.io/github/v/release/ferueda/go-f1-fantasy?sort=semver)](https://github.com/ferueda/go-f1-fantasy/releases)
[![GoDoc](https://godoc.org/github.com/ferueda/go-f1-fantasy?status.svg)](http://godoc.org/github.com/ferueda/go-f1-fantasy)
[![Test Status](https://github.com/ferueda/go-f1-fantasy/workflows/tests/badge.svg)](https://github.com/ferueda/go-f1-fantasy/actions?query=workflow%3Atests)

go-f1-fantasy is a Go client library for accessing the [F1 Fantasy Game API](https://fantasy.formula1.com/).

## Installation

To install the library, simply

`go get github.com/ferueda/go-f1-fantasy`

## Usage
```go
import "github.com/ferueda/go-f1-fantasy/f1"
```

Construct a new client, then use the various services on the client to access different parts of the API. For example:

```go
client := f1.NewClient(nil)
ctx := context.Background()

// list all drivers
drivers, err := client.Drivers.ListAll(ctx)

// get one team by id
team, err := client.Teams.GetOne(ctx, 3)
```

## How to Contribute

* Fork a repository
* Add/Fix something
* Check that tests are passing
* Create PR

Current contributors:

- [Felipe Rueda](https://github.com/ferueda)

## License ##

This library is distributed under the MIT License found in the [LICENSE](./LICENSE)
file.