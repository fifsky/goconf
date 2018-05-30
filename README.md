<p align="center">
<img
    src="logo.png"
    width="240" height="78" border="0" alt="GJSON">
<br>
<a href="https://travis-ci.org/fifsky/goconf"><img src="https://travis-ci.org/fifsky/goconf.svg" alt="Build Status"></a>
<a href="https://codecov.io/gh/fifsky/goconf"><img src="https://codecov.io/gh/fifsky/goconf/branch/master/graph/badge.svg" alt="codecov"></a>
<a href="https://godoc.org/github.com/fifsky/goconf"><img src="https://godoc.org/github.com/gin-gonic/gin?status.svg" alt="GoDoc"></a>
<a href="https://opensource.org/licenses/mit-license.php" rel="nofollow"><img src="https://badges.frapsoft.com/os/mit/mit.svg?v=103"></a>
</p>

<p align="center">Gjson-based configuration file for golang</a></p>

GoConf is a Go package that quickly retrieves JSON configuration files based on [GJSON](https://github.com/tidwall/gjson/)

Getting Started
===============

## Installing

To start using GoConf, install Go and run `go get`:

```sh
$ go get -u github.com/fifsky/goconf
```

This will retrieve the library.


## Get config a value

Config file: testdata/dev.json

```
{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44},
    {"first": "Roger", "last": "Craig", "age": 68},
    {"first": "Jane", "last": "Murphy", "age": 47}
  ]
}
```


```go
package main

import "github.com/fifsky/goconf"

func main() {
    conf, _ := goconf.NewConfig("./testdata/")
    ret, _ := conf.Get("dev.name.last")

    println(ret.String())
}
```

This will print:

```
Anderson
```

You can also use the Must function to simplify this

```go
package main

import "github.com/fifsky/goconf"

func main() {
    conf, _ := goconf.NewConfig("./testdata/")
    println(conf.MustGet("dev.name.last").String())
}
```


> first key `dev` is filename

## Path Syntax
For more Path Syntax, see the [GJSON document](https://github.com/tidwall/gjson/blob/master/README.md#path-syntax)

## Default value
If you use the Must function to access a non-existent configuration file or keys, GoConf returns zero value, which is consistent with gjson.result

```go
conf.MustGet("dev.notfound").String() //value is empty string
conf.MustGet("dev2.not").Int() //value is 0 int
```


## Contact
Xudong Cai [@fifsky](https://fifsky.com/)

## License

GoConf source code is available under the MIT [License](/LICENSE).