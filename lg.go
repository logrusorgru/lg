package lg

import (
	"fmt"
	"log"
	"sync"
)

// Logger is just logger. It writes data to stderr by default.
// You can change its output by log.SetOutput(%io.Writer%)
// It's because this logger based on `log` standart package.
type Logger struct {
	sync.WaitGroup
	in chan string
}

// NewLogger creates a new Logger
func NewLogger() *Logger {
	l := &Logger{}
	l.in = make(chan string, 10)
	l.Add(1)
	go func(in <-chan string, wg *sync.WaitGroup) {
		defer wg.Done()
		for s := range in {
			log.Print(s)
		}
	}(l.in, &l.WaitGroup)
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
	l.Wait()
}
