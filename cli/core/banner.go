package core

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

func ShowBanner() {
	// Print ASCII logo
	figure.NewColorFigure("K8ly", "", "cyan", true).Print()

	// Print subtitle in styled color
	sub := color.New(color.FgHiWhite, color.Bold).SprintFunc()
	link := color.New(color.FgCyan).SprintFunc()

	println(sub("📦 A Cloud-Native DevKit by @Omotolani98"))
	println(link("🌐 https://github.com/Omotolani98/k8ly"))
	println()
}
