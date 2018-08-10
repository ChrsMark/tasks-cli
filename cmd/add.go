package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"strings"
	"github.com/gophersize/task/db"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong:", err)
		}
		fmt.Printf("added [%s] to your list\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}