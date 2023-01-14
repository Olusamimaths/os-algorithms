package memory

import "github.com/olusamimaths/os-algorithms/program"

type RunningJobs []program.Program

type Memory interface {
	GetSize() int
	IsMemoryFree() bool
	GetSizeUsed() int
	LoadJob(job program.Program)
	GetRunningJobs() RunningJobs
}

type memory struct {
	size int
	used int
	isFree bool
	runningJobs RunningJobs
}

func (m *memory) GetSize() int {
	return m.size
}

func (m *memory) IsMemoryFree() bool {
	return m.isFree
}

func (m *memory) GetSizeUsed() int {
	return m.used
}
func (m *memory) GetRunningJobs() RunningJobs {
	return m.runningJobs
}

func(m *memory) LoadJob(job program.Program) {
	jobSize := job.GetSize()
	m.isFree = false
	m.size += jobSize
	m.runningJobs = append(m.runningJobs, job)
}

func NewMainMemory() Memory {
	return &memory{size: 1024, used: 0, isFree: true, runningJobs: RunningJobs{}}
}