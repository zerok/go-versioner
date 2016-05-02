The broad idea here is to provide a little helper tool for gerating a version
info string that can then be used during the build-process of a Go
application.

**WARNING:** Right now this is exclusively a playground for dealing with
versioning and perhaps in the end providing a tool for myself to make this part
of my release-pipeline easier 😊


## Usage

go-versioner supports two kind of output.

First, it can generate a simple Go source file that you can then use for
instance in combination with go-generate to get a version number into your
project:

```
//go:generate go-versioner -output gen_version.go
```

This will generate a new `go_version.go` file that is part of the main package and
that exposes a single constant `VERSION` holding various build information
aspects like the latest found tag name, git-ref and timestamp. For the first
two, `git describe` is used.


The other approach is relying on `go build`'s support for ldflags. First you
need to declare a variable in your main package that can be filled:

```
package main

var VERSION = "unversioned"

// ...
```

Then build the package like this:

```
go build -ldflags "-X main.VERSION='$(go-versioner)'"
```

This is also the approach I've used for this very package. For details take a
look at the Makefile.
