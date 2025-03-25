package model

import (
	"math/rand"
	"time"

	"github.com/charmbracelet/bubbles/table"
)

type Scheduler struct {
	Table          *table.Model
	processQueue   []*Process
	quantum        int
	currentProcess *Process
	stopChan       chan struct{}
}

func NewScheduler(quantum int) *Scheduler {
	Table := NewTableWithStyle()
	Table.SetRows([]table.Row{})

	s := &Scheduler{
		Table:        Table,
		processQueue: make([]*Process, 0),
		quantum:      quantum,
		stopChan:     make(chan struct{}),
	}

	for i := 1; i < 8; i++ {
		s.addProcess(NewProcess(i, rand.Intn(10)+3))
	}
	return s
}

func (s *Scheduler) Start() {
	go func() {
		for {
			select {
			case <-s.stopChan:
				return
			default:
				if !s.run() {
					s.stop()
				}
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	go func() {
		// time.Sleep((rand.Intn(5) + 3) * int(time.Second))
    var sleep time.Duration = time.Duration(rand.Intn(5) + 3)
    time.Sleep(sleep * time.Second)
		for i := 9; i < 16; i++ {
			s.addProcess(NewProcess(i, rand.Intn(10)+3))
		}
		s.Table.SetRows(s.toRows())
	}()
}

func (s *Scheduler) stop() {
	close(s.stopChan)
}

func (s *Scheduler) addProcess(process *Process) {
	s.processQueue = append(s.processQueue, process)
}

func (s *Scheduler) removeProcess(process *Process) {
	for i, p := range s.processQueue {
		if p == process {
			s.processQueue = append(s.processQueue[:i], s.processQueue[i+1:]...)
			break
		}
	}
}

func (s *Scheduler) toRows() []table.Row {
	rows := make([]table.Row, 0)
	for _, process := range s.processQueue {
		rows = append(rows, process.ToRow())
	}
	return rows
}

func (s *Scheduler) run() bool {
	if len(s.processQueue) == 0 {
		return false
	}
	if s.currentProcess == nil {
		s.setNextProcess()
		return true
	} else {
		s.runProcess()
		return true
	}
}

func (s *Scheduler) setNextProcess() {
	if s.currentProcess != nil {
		return
	}

	for _, process := range s.processQueue {
		if process.State == New {
			process.State = Ready
		}
	}

	maxWaitTime := 0
	s.currentProcess = s.processQueue[0]
	for i, process := range s.processQueue {
		if process.TimeInQueue > maxWaitTime {
			maxWaitTime = process.TimeInQueue
			s.currentProcess = process
			s.Table.SetCursor(i)
		}
	}

	s.currentProcess.State = Running
	s.Table.SetRows(s.toRows())
}

func (s *Scheduler) runProcess() {
	if s.currentProcess == nil {
		return
	}

	minTime := min(s.currentProcess.TimeRemaining, s.quantum)
	for _, process := range s.processQueue {
			process.TimeInQueue++
	}

	removed := false
	for i := 0; i < minTime; i++ {
		time.Sleep(time.Second)
		s.currentProcess.TimeRemaining--
		s.Table.SetRows(s.toRows())
		if s.currentProcess.TimeRemaining == 0 {
			s.removeProcess(s.currentProcess)
			removed = true
			break
		}
	}

	if !removed {
		s.currentProcess.State = Ready

    current := s.currentProcess
    s.removeProcess(current)
    s.addProcess(current)
	}
	s.currentProcess = nil
	s.Table.SetRows(s.toRows())
}

func min(timeRemaining int, quantum int) int {
	if timeRemaining < quantum {
		return timeRemaining
	}
	return quantum
}
