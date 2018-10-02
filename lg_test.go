//
// Copyright (c) 2015 Konstantin Ivanov <kostyarin.ivanov@gmail.com>.
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

package lg

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"reflect"
	"sync"
	"testing"
)

var logLock = new(sync.Mutex)

type Capture struct {
	stderr *os.File
	buffer *bytes.Buffer
	lines  []string
}

func Cap() *Capture {
	c := new(Capture)
	c.buffer = new(bytes.Buffer)
	logLock.Lock()
	log.SetOutput(c.buffer)
	return c
}

func (c *Capture) Release() {
	log.SetOutput(os.Stderr) // restore
	logLock.Unlock()
	for {
		l, err := c.buffer.ReadString('\n')
		if err != nil {
			break
		}
		if len(l) > 21 /*&& l[len(l)-1] == '\n'*/ {
			// 2016/01/13 01:38:39 <--- cut this --->\n
			c.lines = append(c.lines, l[20:]) /*len(l)-1])*/
		}
	}
	c.stderr = nil
	c.buffer = nil
}

func (c *Capture) CompareAndRemove(s string) (ok bool) {
	for i, v := range c.lines {
		if v == s {
			ok = true
			c.lines[i] = c.lines[len(c.lines)-1]
			c.lines = c.lines[:len(c.lines)-1]
			return
		}
	}
	return
}
func (c *Capture) Remains() int {
	return len(c.lines)
}

type pp interface {
	Println(err ...interface{})
	Printf(format string, args ...interface{})
}

type s interface {
	Stop()
}

func TestNewLogger(t *testing.T) {
	l := NewLogger()
	defer l.Stop()
	if l == nil {
		t.Error("NewLogger returns a nil value")
		t.FailNow()
	}
	if !reflect.TypeOf(l).Implements(reflect.TypeOf((*pp)(nil)).Elem()) {
		t.Error("Logger doesn't implements Println and Printf")
	}
	if !reflect.TypeOf(l).Implements(reflect.TypeOf((*s)(nil)).Elem()) {
		t.Error("Logger hasn't a Stop method")
	}
	if cap(l.in) != defaultBufferSize {
		t.Error("wrong buffer size for default logger")
	}
}

func TestNewLoggerBufer(t *testing.T) {
	bSize := 0
	l := NewLoggerBuffer(bSize)
	defer l.Stop()
	if l == nil {
		t.Error("NewLogger returns a nil value")
		t.FailNow()
	}
	if !reflect.TypeOf(l).Implements(reflect.TypeOf((*pp)(nil)).Elem()) {
		t.Error("Logger doesn't implements Println and Printf")
	}
	if !reflect.TypeOf(l).Implements(reflect.TypeOf((*s)(nil)).Elem()) {
		t.Error("Logger hasn't a Stop method")
	}
	if cap(l.in) != bSize {
		t.Errorf("wrong buffer size, expected %d, got %d",
			bSize, cap(l.in))
	}
}

const GoLimit = 10

func TestLogger_Println(t *testing.T) {
	c := Cap()
	l := NewLogger()
	wg := new(sync.WaitGroup)
	wg.Add(GoLimit)
	for i := 0; i < GoLimit; i++ {
		go func(n int) {
			defer wg.Done()
			l.Println("gorutine no", n)
		}(i)
	}
	l.Println("waiting...")
	wg.Wait()
	l.Println("stopping")
	l.Stop()
	l.Println("will never printed")
	c.Release()
	lines := make([]string, 0, GoLimit+2)
	for i := 0; i < GoLimit; i++ {
		lines = append(lines, fmt.Sprintln("gorutine no", i))
	}
	lines = append(lines, "waiting...\n")
	lines = append(lines, "stopping\n")
	if c.Remains() != len(lines) {
		t.Errorf("wrong lines number, expected %d, got %d",
			len(lines), c.Remains())
		t.FailNow()
	}
	for _, s := range lines {
		if !c.CompareAndRemove(s) {
			t.Error("bad output for:", s)
		}
	}
}

func TestLogger_Printf(t *testing.T) {
	c := Cap()
	l := NewLogger()
	wg := new(sync.WaitGroup)
	wg.Add(GoLimit)
	for i := 0; i < GoLimit; i++ {
		go func(n int) {
			defer wg.Done()
			l.Printf("gorutine no '%0.4d'", n)
		}(i)
	}
	l.Println("waiting...")
	wg.Wait()
	l.Println("stopping")
	l.Stop()
	l.Println("will never printed")
	c.Release()
	lines := make([]string, 0, GoLimit+2)
	for i := 0; i < GoLimit; i++ {
		lines = append(lines, fmt.Sprintf("gorutine no '%0.4d'\n", i))
	}
	lines = append(lines, "waiting...\n")
	lines = append(lines, "stopping\n")
	if c.Remains() != len(lines) {
		t.Errorf("wrong lines number, expected %d, got %d",
			len(lines), c.Remains())
		t.FailNow()
	}
	for _, s := range lines {
		if !c.CompareAndRemove(s) {
			t.Error("bad output for:", s)
		}
	}
}

func TestLogger_Stop(t *testing.T) {
	c := Cap()
	l := NewLogger()
	l.Stop()
	l.Println("one")
	l.Println("two")
	l.Println("three")
	l.Println("will never printed")
	c.Release()
	if c.Remains() != 0 {
		t.Errorf("no empty output, got %d lines", c.Remains())
	}
}
