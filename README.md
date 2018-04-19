lg
==

[![GoDoc](https://godoc.org/github.com/logrusorgru/lg?status.svg)](https://godoc.org/github.com/logrusorgru/lg)
[![WTFPL License](https://img.shields.io/badge/license-wtfpl-blue.svg)](http://www.wtfpl.net/about/)
[![Build Status](https://travis-ci.org/logrusorgru/lg.svg)](https://travis-ci.org/logrusorgru/lg)
[![Coverage Status](https://coveralls.io/repos/logrusorgru/lg/badge.svg?branch=master)](https://coveralls.io/r/logrusorgru/lg?branch=master)
[![GoReportCard](https://goreportcard.com/badge/logrusorgru/lg)](https://goreportcard.com/report/logrusorgru/lg)
[![Gitter](https://img.shields.io/badge/chat-on_gitter-46bc99.svg?logo=data:image%2Fsvg%2Bxml%3Bbase64%2CPHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIGhlaWdodD0iMTQiIHdpZHRoPSIxNCI%2BPGcgZmlsbD0iI2ZmZiI%2BPHJlY3QgeD0iMCIgeT0iMyIgd2lkdGg9IjEiIGhlaWdodD0iNSIvPjxyZWN0IHg9IjIiIHk9IjQiIHdpZHRoPSIxIiBoZWlnaHQ9IjciLz48cmVjdCB4PSI0IiB5PSI0IiB3aWR0aD0iMSIgaGVpZ2h0PSI3Ii8%2BPHJlY3QgeD0iNiIgeT0iNCIgd2lkdGg9IjEiIGhlaWdodD0iNCIvPjwvZz48L3N2Zz4%3D&logoWidth=10)](https://gitter.im/logrusorgru/lg?utm_source=share-link&utm_medium=link&utm_campaign=share-link)

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

It's just package for demonstration. But you can use it in production.
It's impossible to use `log.Lshortfile` and `log.Llongfile` flags in
this logger.

# License

Copyright © 2015 Konstantin Ivanov <kostyarin.ivanov@gmail.com>  
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE file for more details.
