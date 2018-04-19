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

package lg

import (
	"log"
	"os"
)

func ExampleLogger_Println() {
	// set output to stdout
	// intead of stderr (default)
	log.SetOutput(os.Stdout)
	// omit date-time (for example)
	log.SetFlags(0)
	// create new logger
	l := NewLogger()
	// stop it after
	defer l.Stop()
	l.Println("Hello, logger!")
	// Output:
	// Hello, logger!
}

func ExampleLogger_Printf() {
	// set output to stdout
	// intead of stderr (default)
	log.SetOutput(os.Stdout)
	// omit date-time (for example)
	log.SetFlags(0)
	// create new logger
	l := NewLogger()
	// stop it after
	defer l.Stop()
	l.Printf("Hello, %d times!", 909)
	// Output:
	// Hello, 909 times!
}
