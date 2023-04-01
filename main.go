package main

import (
	"github.com/spf13/cobra"
	"github.com/tiptophelmet/mywireguard/cmd"
)

func main() {
	cobra.OnInitialize(validateDependencies)
	cmd.Execute()
}
