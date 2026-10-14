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

	createCmd := &cobra.Command{
		Use:  "create",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Generating K8s-Manifest for %s\n", args[0])
		},
	}
	rootCmd.AddCommand(createCmd)
	rootCmd.Execute()
}
