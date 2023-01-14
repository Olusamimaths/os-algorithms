package program

type Program interface {
	Run()
	GetSize() int
	GetId() string
	GetInstructions() InstructionsType
}

type InstructionsType []Instruction

type program struct {
	id string
	size int
	instructions InstructionsType
}

func (p *program) Run() {
	print("Running program ", p.id, " with size ", p.size)
}
func (p *program) GetSize() int {
	return p.size
}
func (p *program) GetId() string {
	return p.id
}
func (p *program) GetInstructions() InstructionsType {
	return p.instructions
}

type CreateNewProgramInput struct {
	Id string
	Size int
	Instructions InstructionsType
}

func NewProgram(input CreateNewProgramInput) Program{
	return &program{size: input.Size, id: input.Id, instructions: input.Instructions }
}

