//
// Copyright (c) 2015 Konstanin Ivanov <kostyarin.ivanov@gmail.com>.
// All rights reserved. This program is free software. It comes without
// any warranty, to the extent permitted by applicable law. You can
// redistribute it and/or modify it under the terms of the Do What
// The Fuck You Want To Public License, Version 2, as published by
// Sam Hocevar. See LICENSE file for more details or see below.
//

//
//        DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//                    Version 2, December 2004
//
// Copyright (C) 2004 Sam Hocevar <sam@hocevar.net>
//
// Everyone is permitted to copy and distribute verbatim or modified
// copies of this license document, and changing it is allowed as long
// as the name is changed.
//
//            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
//  0. You just DO WHAT THE FUCK YOU WANT TO.
//

// Package lg is just primitive logger for Golang applications,
// yeah it's thread safe. But it's useless if you want to use
// log.Shortfile or log.Longfile log flags.
package lg

import (
	"fmt"
	"log"
	"sync"
)

const defaultBufferSize = 10

// Logger is just logger. It writes data to stderr by default.
// You can change its output by log.SetOutput(%io.Writer%)
// It's because this logger based on `log` standart package.
type Logger struct {
	in chan string
	wg sync.WaitGroup
}

// NewLogger creates a new Logger with default buffer size
func NewLogger() *Logger {
	return NewLoggerBuffer(defaultBufferSize)
}

// NewLoggerBuffer creates a new Logger with provided buffer size.
// The buffer is just buffer of internal channel.
func NewLoggerBuffer(n int) *Logger {
	l := &Logger{
		in: make(chan string, n),
		wg: sync.WaitGroup{},
	}
	l.wg.Add(1)
	go func(in <-chan string, wg *sync.WaitGroup) {
		defer wg.Done()
		for s := range in {
			log.Print(s)
		}
	}(l.in, &l.wg)
	return l

}

// Println works like `log.Println`
func (l *Logger) Println(err ...interface{}) {
	defer func() { recover() }()
	l.in <- fmt.Sprintln(err...)
}

// Printf works like `log.Printf`
func (l *Logger) Printf(format string, args ...interface{}) {
	defer func() { recover() }()
	l.in <- fmt.Sprintf(format, args...)
}

// Stop stops the logger. After that, the logger nothing will print.
func (l *Logger) Stop() {
	close(l.in)
	l.wg.Wait()
}
