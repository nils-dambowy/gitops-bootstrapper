package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "Bootstrapper",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(os.Stdout, "Bootstrapper starting...")
		},
	}

	rootCmd.Execute()
}
