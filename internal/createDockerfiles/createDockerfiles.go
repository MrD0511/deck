package createDockerfiles

import (
	"fmt"
	"github.com/MrD0511/deck/internal/stack"
)

func MainFunc() {
	report, err := stack.DetectFramework()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(report)
	stack.PrintTechStackReport(report)
}
