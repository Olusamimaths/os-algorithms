package memory

import "github.com/olusamimaths/os-algorithms/program"

type Memory interface {
	LoadJob(job program.Program)
}

type memory struct {
	Size int
	Used int
	IsFree bool
	RunningJobs []program.Program
}

func(m *memory) LoadJob(job program.Program) {
	jobSize := job.GetSize()
	m.IsFree = false
	m.Size += jobSize
	m.RunningJobs = append(m.RunningJobs, job)
}