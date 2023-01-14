package memory

import "github.com/olusamimaths/os-algorithms/program"

type Memory interface {
	GetSize() int
	IsMemoryFree() bool
	GetSizeUsed() int
	LoadJob(job program.Program)
}

type memory struct {
	Size int
	Used int
	IsFree bool
	RunningJobs []program.Program
}

func (m *memory) GetSize() int {
	return m.Size
}

func (m *memory) IsMemoryFree() bool {
	return m.IsFree
}

func (m *memory) GetSizeUsed() int {
	return m.Used
}

func(m *memory) LoadJob(job program.Program) {
	jobSize := job.GetSize()
	m.IsFree = false
	m.Size += jobSize
	m.RunningJobs = append(m.RunningJobs, job)
}

func NewMainMemory() Memory {
	return &memory{Size: 1024, Used: 0, IsFree: true, RunningJobs: []program.Program{}}
}