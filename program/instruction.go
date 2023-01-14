package program

type Instruction interface {
	GetSize() int
	GetAction() string
	Execute()
}
type instruction struct {
	Size int
	Action string
}
func (i * instruction) GetSize() int {
	return i.Size
}
func (i * instruction) GetAction() string {
	return i.Action
}
func (i *instruction) Execute() {
	print("Executing instruction ", i.Action, " with Size ", i.Size)
}

type CreateNewInstructionInput struct {
	Size int
	Action string
}

func NewInstruction(input CreateNewInstructionInput) Instruction {
	return &instruction{
		Size: input.Size,
		Action: input.Action,
	}
}
