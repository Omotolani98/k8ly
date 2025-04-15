package main

import "github.com/Omotolani98/k8ly/cli/cmd"

var version = "dev"

func main() {
  cmd.Version = version
  cmd.Execute()
}
