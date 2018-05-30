<p align="center">
<img
    src="logo.png"
    width="240" height="78" border="0" alt="GJSON">
<br>
<a href="https://travis-ci.org/fifsky/goconf"><img src="https://img.shields.io/travis/fifsky/goconf.svg?style=flat-square" alt="Build Status"></a>
<a href="https://github.com/fifsky/goconf" rel="nofollow"><img src="https://camo.githubusercontent.com/c119511be2f77b84fe0d0df8621f32971c239d70/68747470733a2f2f706f7365722e707567782e6f72672f766572792f6672616d65776f726b2f6c6963656e73652e737667" alt="License" data-canonical-src="https://poser.pugx.org/very/framework/license.svg" style="max-width:100%;"></a>
</p>

<p align="center">Gjson-based configuration file</a></p>

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
If you use the Must function to access a non-existent configuration file or keys, GoConf returns zore value, which is consistent with gjson.result

```go
conf.MustGet("dev.notfound").String() //value is empty string
conf.MustGet("dev2.not").Int() //value is 0 int
```


## Contact
Xudong Cai [@fifsky](https://fifsky.com/)

## License

GoConf source code is available under the MIT [License](/LICENSE).