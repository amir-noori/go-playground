package patterns

import (
	"fmt"
	"time"
)

type Process interface {
	run()
	getProcessId() int
}

type LightProcess struct {
	processId int
}

func (lp *LightProcess) run() {
	fmt.Println("running light process")
	counter := 0
	for i := 0; i < 10000000; i++ {
		counter++
	}
}

func (lp *LightProcess) getProcessId() int {
	return lp.processId
}

type logger struct {
	process Process
}

func (l *logger) run() {
	fmt.Printf("process id: %d\n", l.process.getProcessId())
	l.process.run()
}

func (l *logger) getProcessId() int {
	return l.process.getProcessId()
}

type timer struct {
	process Process
}

func (t *timer) run() {
	t1 := time.Now()
	t.process.run()
	t2 := time.Now()
	diff := t2.Sub(t1)
	fmt.Printf("time consumed %s\n", diff)
}

func (t *timer) getProcessId() int {
	return t.process.getProcessId()
}

func RunTypeDecorator() {
	process1 := &LightProcess{
		processId: 100,
	}
	timedProcess1 := &timer{
		process1,
	}
	loggedTimedProcess1 := &logger{
		timedProcess1,
	}

	loggedTimedProcess1.run()
}
