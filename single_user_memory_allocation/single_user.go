package single_user

import (
	"github.com/olusamimaths/os-algorithms/memory"
	prog "github.com/olusamimaths/os-algorithms/program"
)

func CreateSingleUserProgram() prog.Program {
	inst1 := prog.NewInstruction(prog.CreateNewInstructionInput{
		Size:   5,
		Action: "mov   ax,   8h ",
	})

	inst2 := prog.NewInstruction(prog.CreateNewInstructionInput{
		Size:   10,
		Action: "and   ax, 1",
	})

	inst3 := prog.NewInstruction(prog.CreateNewInstructionInput{
		Size:   5,
		Action: "int   0x80 ",
	})

	inst4 := prog.NewInstruction(prog.CreateNewInstructionInput{
		Size: 30,
		Action: `mov   eax, 4  
				 mov   ebx, 1  
				 mov   ecx, odd_msg `,
	})

	program := prog.NewProgram(prog.CreateNewProgramInput{
		Id:           "Job 1",
		Size:         50,
		Instructions: []prog.Instruction{inst1, inst2, inst3, inst4},
	})
	return program
	
}

func LoadJobsToMemory() {
	program := CreateSingleUserProgram()
	mainMemory := memory.NewMainMemory()

	programSize := program.GetSize()
	mainMemorySize := mainMemory.GetSize()

	if programSize > mainMemorySize {
		print("Can't load program ", program.GetSize())
	}
	mainMemory.LoadJob(program)
}
