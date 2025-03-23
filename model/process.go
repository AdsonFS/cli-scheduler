package model

import "fmt"

const (
	New   = "new"
	Ready = "ready"
	// Waiting    = "waiting"
	Running    = "running"
	Terminated = "terminated"
)

type Process struct {
	PID           int
	TimeInQueue   int
	TimeRemaining int
	State         string
}

func NewProcess(pid, totalTime int) *Process {
	return &Process{
		PID:           pid,
		TimeInQueue:   0,
		TimeRemaining: totalTime,
		State:         New,
	}
}

func (p *Process) ToRow() []string {
  return []string{
    fmt.Sprintf("%d", p.PID),
    fmt.Sprintf("%ds", p.TimeInQueue),
    fmt.Sprintf("%ds", p.TimeRemaining),
    p.State,
  }
}
