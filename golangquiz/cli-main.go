/*
Copyright © 2022 Romina Lodolo <lodoloromina@gmail.com>
*/
package main

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
)

func main() {
	cmdPrint := &cobra.Command{
		Use:   "question [ID]",
		Short: "Fetch question",
		Long:  `Fetch question`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("question: " + strings.Join(args, " "))
		},
	}

	rootCmd := &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdPrint)
	rootCmd.Execute()
}
