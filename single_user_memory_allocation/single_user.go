package single_user

import prog "github.com/olusamimaths/os-algorithms/program"

func LoadJobs() {
	inst1 := prog.NewInstruction(prog.CreateNewInstructionInput{
		Size:   5,
		Action: "LOAD AX into BX",
	})

	inst2 := prog.NewInstruction(prog.CreateNewInstructionInput{
		Size:   5,
		Action: "LOAD AX into BX",
	})

	program1 := prog.NewProgram(prog.CreateNewProgramInput{
		Id:           "Job 1",
		Size:         50,
		Instructions: []prog.Instruction{inst1, inst2},
	})
}
