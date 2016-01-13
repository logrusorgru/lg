package lg

import (
	"fmt"
	"log"
	"sync"
)

type Logger struct {
	sync.WaitGroup
	in chan string
}

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

func (l *Logger) Println(err ...interface{}) {
	defer func() { recover() }()
	l.in <- fmt.Sprintln(err...)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	defer func() { recover() }()
	l.in <- fmt.Sprintf(format, args...)
}

func (l *Logger) Stop() {
	close(l.in)
	l.Wait()
}
