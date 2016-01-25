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
