/*
Copyright © 2019 atselvan

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Addition of two numbers",
	Long: `The command adds two numbers
	result = num1 + num2
`,
	Run: func(cmd *cobra.Command, args []string) {
		num1, _ := cmd.Flags().GetInt32("num1")
		num2, _ := cmd.Flags().GetInt32("num2")
		fmt.Println(num1 + num2)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().Int32P("number1", "n", 0, "number 1")
	addCmd.Flags().Int32P("number2", "m", 0, "number 2")
	addCmd.Flags().StringP("test", "s", "", "")
	addCmd.Flags().StringP("test1", "p", "", "")
	addCmd.MarkFlagRequired("test")
	addCmd.MarkFlagRequired("test1")
}
