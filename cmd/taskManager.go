package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addTaskCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your application",
	Long:  `A Longer description`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := NewTM.CreateTask(args[0], false)
		fmt.Println(task)
	},
}

var updateTaskCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your application",
	Long:  `A Longer description`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		updatedTask := NewTM.UpdateTask(args[0])
		fmt.Println(updatedTask)
	},
}

var deleteTaskCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your application",
	Long:  `A Longer description`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		deletedTask := NewTM.DeleteTask(args[0])
		fmt.Println(deletedTask)
	},
}

func init() {

	rootCmd.AddCommand(addTaskCmd)
	rootCmd.AddCommand(updateTaskCmd)
	rootCmd.AddCommand(deleteTaskCmd)
}
