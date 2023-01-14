package program

type Program interface {
	Run()
	GetSize() int
	GetInstructions() InstructionsType
}

type InstructionsType []Instruction

type program struct {
	Id string
	Size int
	Instructions InstructionsType
}

func (p *program) Run() {
	print("Running program ", p.Id, " with Size ", p.Size)
}
func (p *program) GetSize() int {
	return p.Size
}
func (p *program) GetInstructions() InstructionsType {
	return p.Instructions
}

type CreateNewProgramInput struct {
	Id string
	Size int
	Instructions InstructionsType
}

func NewProgram(input CreateNewProgramInput) Program{
	return &program{Size: input.Size, Id: input.Id, Instructions: input.Instructions }
}

