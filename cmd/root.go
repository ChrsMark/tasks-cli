package cmd

import (
	"github.com/spf13/cobra"
	)

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI, for task management",
	//Run: func(cmd *cobra.Command, args []string) {
	//	// Do Stuff Here
	//},
}

//func Execute() {
//	if err := rootCmd.Execute(); err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//}
