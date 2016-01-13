lg
==

[![GoDoc](https://godoc.org/github.com/logrusorgru/lg?status.svg)](https://godoc.org/github.com/logrusorgru/lg)
[![WTFPL License](https://img.shields.io/badge/license-wtfpl-blue.svg)](http://www.wtfpl.net/about/)
[![Build Status](https://travis-ci.org/logrusorgru/lg.svg)](https://travis-ci.org/logrusorgru/lg)
[![Coverage Status](https://coveralls.io/repos/logrusorgru/lg/badge.svg?branch=master)](https://coveralls.io/r/logrusorgru/lg?branch=master)
[![GoReportCard](http://goreportcard.com/badge/logrusorgru/lg)](http://goreportcard.com/report/logrusorgru/lg)

it's just primitive logger for Golang applications, yeah it's thread safe

# Available methods

- `Printf(format string, args ...interface{})`
- `Println(err ...interface{})`

# How to use

Get it

```bash
go get github.com/logrusorgru/lg
```

Test it

```bash
go test github.com/logrusorgru/lg
```

Use it

```go
package main

import(
	"github.com/logrusorgru/lg"
)

// Create a logger
var l = lg.NewLogger()

func main() {
	// stop it
	defer l.Stop()
	c1, c2 := make(chan struct{}), make(chan struct{})
	go func(){
		l.Printf("hello from gorutine number %d", 1)
		c1 <- struct{}{} // done
	}()
	go func(){
		l.Printf("hello from gorutine number %d", 2)
		c2 <- struct{}{} // done
	}()
	<-c1
	<-c2
	l.Println("Done")
}
```

# Note

It's just package for demonstration. But you can to use it in production.

# License

Copyright © 2015 Konstantin Ivanov <ivanov.konstantin@logrus.org.ru>  
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE.md file for more details.
