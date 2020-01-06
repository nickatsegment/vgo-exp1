package main

import (
	"fmt"
	"os"

	"github.com/nickatsegment/vgo-exp1/lib/v2"
	"github.com/spf13/cobra"
)

var Version = "2.0.0-rc1.0"

var rootCmd = &cobra.Command{
	Use: "vgo-exp1",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("cli version: %s\n", Version)
		fmt.Printf("lib version: %s\n", lib.Version())
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
