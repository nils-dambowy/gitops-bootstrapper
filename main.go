package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	yamlTemplate := `	apiVersion: apps/v1
kind: Deployment
metadata:
  name: %s
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: %s
        image: nginx:1.14.2
        ports:
        - containerPort: 80

	`

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
			x := fmt.Sprintf(yamlTemplate, args[0], args[0])
			fmt.Println(x)
		},
	}
	rootCmd.AddCommand(createCmd)
	rootCmd.Execute()
}
