package core

import (
	"fmt"
	"github.com/fatih/color"
)

var (
	Info    = color.New(color.FgCyan).SprintFunc()
	Success = color.New(color.FgGreen).SprintFunc()
	Error   = color.New(color.FgRed).SprintFunc()
	Section = color.New(color.FgHiWhite, color.Bold).SprintFunc()
)

func PrintSection(title string) {
	fmt.Println(Section("\n" + title))
	fmt.Println(Section("────────────────────────────────"))
}

func PrintInfo(msg string) {
	fmt.Println(Info("📎 " + msg))
}

func PrintSuccess(msg string) {
	fmt.Println(Success("✅ " + msg))
}

func PrintError(msg string) {
	fmt.Println(Error("❌ " + msg))
}
