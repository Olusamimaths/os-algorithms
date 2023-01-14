package program

type Program interface {
	Run()
}

type program struct {
	Id string
	Size int
	Instructions []instruction
}

func (p *program) Run() {
	print("Running program ", p.Id, " with Size ", p.Size)
}

type CreateNewProgramInput struct {
	program
}

func NewProgram(input CreateNewProgramInput) program{
	return program{Size: input.Size, Id: input.Id, Instructions: input.Instructions }
}