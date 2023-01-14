package program

type Instruction interface {
	Execute()
}
type instruction struct {
	size int
	action string
}
func (i *instruction) Execute() {
	print("Executing instruction ", i.action, " with Size ", i.size)
}

