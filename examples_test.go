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
	log.SetFlags(log.Lshortfile)
	// create new logger
	l := NewLogger()
	// stop it after
	defer l.Stop()
	l.Println("Hello, logger!")
	// Output:
	// lg.go:37: Hello, logger!
}

func ExampleLogger_Printlf() {
	// set output to stdout
	// intead of stderr (default)
	log.SetOutput(os.Stdout)
	// omit date-time (for example)
	log.SetFlags(log.Lshortfile)
	// create new logger
	l := NewLogger()
	// stop it after
	defer l.Stop()
	l.Printf("Hello, %d times!", 909)
	// Output:
	// lg.go:37: Hello, 909 times!
}
