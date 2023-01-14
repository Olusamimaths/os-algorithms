package program

type Program interface {
	Run()
	GetSize() int
	GetId() string
	GetInstructions() InstructionsType
	
	computeProgramSize() int
}

type InstructionsType []Instruction

type program struct {
	id string
	instructions InstructionsType
}

func (p *program) Run() {
	print("Running program ", p.id, " with size ", p.GetSize())
}
func (p *program) GetSize() int {
	return p.computeProgramSize()
}
func (p *program) GetId() string {
	return p.id
}
func (p *program) GetInstructions() InstructionsType {
	return p.instructions
}
func(p * program) computeProgramSize() int {
	sum := 0
	if len(p.instructions) == 0 {
		return sum
	}
	for _, inst := range p.instructions {
		sum += inst.GetSize()
	}
	return sum
}

type CreateNewProgramInput struct {
	Id string
	Instructions InstructionsType
}

func NewProgram(input CreateNewProgramInput) Program{
	return &program{ id: input.Id, instructions: input.Instructions }
}

