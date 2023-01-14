package main

import (
	"fmt"

	single_user "github.com/olusamimaths/os-algorithms/single_user_memory_allocation"
)

func main() {
	fmt.Print("Single User Memory Allocation\n")
	single_user.LoadJobsToMemory()
}